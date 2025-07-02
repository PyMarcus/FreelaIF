import pika
import json
import logging
from config import get_config
from email_sender import send_email
import time

logging.basicConfig(level=logging.INFO)

# Aguarda config-server
config = get_config()

# Conexão com RabbitMQ
while True:
    try:
        params = pika.URLParameters(config['rabbitmq_url'])
        connection = pika.BlockingConnection(params)
        channel = connection.channel()
        channel.queue_declare(queue=config['email_queue'], durable=True)
        break
    except Exception as e:
        logging.warning(f'RabbitMQ não disponível: {e}')
        time.sleep(3)

def callback(ch, method, properties, body):
    try:
        data = json.loads(body)
        subject = data['subject']
        email = data['email']
        name = data['nome']
        content_html = data['conteudo_html']
        send_email(subject, email, name, content_html)
        logging.info(f'Email sent to {email}')
    except Exception as e:
        logging.error(f'Erro ao enviar e-mail: {e}')
    ch.basic_ack(delivery_tag=method.delivery_tag)

channel.basic_qos(prefetch_count=1)
channel.basic_consume(queue=config['email_queue'], on_message_callback=callback)

logging.info('Aguardando mensagens em email.send...')
channel.start_consuming()
