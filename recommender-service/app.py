
import pika
import json
import logging
from config import get_config
from database.db import get_engine as get_rec_engine, create_tables as create_rec_tables, get_session as get_rec_session, Recommendation
from database.user_db import get_engine as get_user_engine, get_session as get_user_session, get_user_model
from services.recommender import generate_recommendations
import time

logging.basicConfig(level=logging.INFO)

# Aguarda config-server
config = get_config()

# Conexão com o banco de recomendações
rec_engine = get_rec_engine(config['db_url'])
create_rec_tables(rec_engine)

# Conexão com o banco de usuários 
user_db_url = config.get('user_db_url', config['db_url'])
user_table = config.get('user_table', 'users')
user_engine = get_user_engine(user_db_url)
create_user_tables(user_engine, user_table)
User = get_user_model(user_table)

# Conexão com RabbitMQ
while True:
    try:
        params = pika.URLParameters(config['rabbitmq_url'])
        connection = pika.BlockingConnection(params)
        channel = connection.channel()
        channel.queue_declare(queue=config['user_queue'], durable=True)
        channel.queue_declare(queue=config['email_queue'], durable=True)
        break
    except Exception as e:
        logging.warning(f'RabbitMQ não disponível: {e}')
        time.sleep(3)

def callback(ch, method, properties, body):
    try:
        data = json.loads(body)
        user_id = data['user_id']
        # Fetch user from user DB
        user_session = get_user_session(user_engine)
        user = user_session.query(User).filter_by(user_id=user_id).first()
        if not user:
            logging.error(f'User {user_id} not found in table {user_table}')
            user_session.close()
            ch.basic_ack(delivery_tag=method.delivery_tag)
            return
        email = user.email
        name = user.nome
        user_session.close()
        recommendations = generate_recommendations()
        # Save to recommendations DB
        rec_session = get_rec_session(rec_engine)
        rec = Recommendation(user_id=user_id, email=email, nome=name, recomendacoes=recommendations)
        rec_session.add(rec)
        rec_session.commit()
        rec_session.close()
        # Generate recommendations HTML (only the div)
        html_div = '''
        <div style="display: flex; justify-content: space-between; flex-wrap: wrap; gap: 20px; margin-top: 30px; margin-bottom: 30px;">
        '''

        for r in recommendations:
            html_div += f'''
            <div style="background-color: #f9f9f9; border-radius: 8px; padding: 20px; width: calc(33.33% - 14px); min-width: 160px; box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);">
            <div style="font-weight: 600; font-size: 16px; margin-bottom: 10px; color: #444444;">{r['title']}</div>
            <div style="font-size: 14px; color: #666666; margin-bottom: 15px;">{r['description']}</div>
            <a href="{r['link']}" style="display: inline-block; background-color: #555555; color: white; padding: 8px 16px; border-radius: 4px; text-decoration: none; font-size: 14px; font-weight: 500;">Acessar projeto</a>
            </div>
            '''

        html_div += '</div>'

        # Publish to email queue
        msg = json.dumps({
            "subject": "FreelaIF - Recomendações", 
            "email": email,
            "nome": name,
            "conteudo_html": html_div
        })
        channel.basic_publish(
            exchange='',
            routing_key=config['email_queue'],
            body=msg,
            properties=pika.BasicProperties(delivery_mode=2)
        )
        logging.info(f'Recommendations generated and published for {email}')
    except Exception as e:
        logging.error(f'Erro no processamento: {e}')
    ch.basic_ack(delivery_tag=method.delivery_tag)

channel.basic_qos(prefetch_count=1)
channel.basic_consume(queue=config['user_queue'], on_message_callback=callback)

logging.info('Aguardando mensagens em user.created...')
channel.start_consuming()
