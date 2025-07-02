import logging
from email.message import EmailMessage

import smtplib
from email.message import EmailMessage
import logging
from config import get_config

config = get_config()


def send_email(subject, email, name, content_html):
    html = f'''
    <html>
    <body style="margin: 0; padding: 0; box-sizing: border-box; font-family: 'Poppins', sans-serif; background-color: #f5f5f5; color: #333333; line-height: 1.6;">
  <div style="max-width: 600px; margin: 20px auto; background-color: #ffffff; border-radius: 8px; overflow: hidden; box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);">

    <!-- Header -->
    <div style="background-color: #333333; padding: 25px 30px; text-align: center;">
      <div style="font-size: 28px; font-weight: 700; color: #ffffff; letter-spacing: 1px;">
        Freela<span style="color: #a3a3a3;">IF</span>
      </div>
    </div>

    <!-- Body -->
    <div style="padding: 30px;">
      <div style="font-size: 20px; margin-bottom: 25px; color: #333333; font-weight: 500;">
        Olá, <b>{name}</b>
      </div>

      <p style="margin-bottom: 20px;">Esperamos que esteja tudo bem com você. Aqui estão seus projetos atuais:</p>

       {content_html}

      <p style="margin-bottom: 20px;">Não hesite em entrar em contato se precisar de qualquer assistência.</p>
      <p>Atenciosamente,<br>Equipe FreelaIF</p>
    </div>

    <!-- Footer -->
    <div style="background-color: #f1f1f1; padding: 20px 30px; text-align: center; font-size: 12px; color: #777777; border-top: 1px solid #e0e0e0;">
      <p style="margin-bottom: 8px;">© 2025 FreelaIF. Todos os direitos reservados.</p>
      <p style="margin-bottom: 8px;">Este email foi enviado para {email} porque você está cadastrado em nossa plataforma.</p>
      <p style="margin-bottom: 8px;">Se você não deseja mais receber nossos emails, <a href="#" style="color: #555555; text-decoration: none;">clique aqui para cancelar a inscrição</a>.</p>

      <div style="margin-top: 12px;">
        <a href="#" style="color: #555555; margin: 0 8px; text-decoration: none;">Termos de Uso</a>
        <a href="#" style="color: #555555; margin: 0 8px; text-decoration: none;">Política de Privacidade</a>
        <a href="#" style="color: #555555; margin: 0 8px; text-decoration: none;">Contato</a>
      </div>

      <p style="margin-top: 12px;">FreelaIF - Conectando talentos e oportunidades</p>
    </div>

  </div>
</body>
    </html>
    '''
    msg = EmailMessage()
    msg['Subject'] = subject
    msg['From'] = config.get('smtp_user', 'noreply@example.com')
    msg['To'] = email
    msg.add_alternative(html, subtype='html')
    try:
        with smtplib.SMTP(config['smtp_host'], config['smtp_port']) as server:
            server.starttls()
            server.login(config['smtp_user'], config['smtp_pass'])
            server.send_message(msg)
        logging.info(f"Email sent to {email} ({name})")
    except Exception as e:
        logging.error(f"Failed to send email to {email}: {e}")
