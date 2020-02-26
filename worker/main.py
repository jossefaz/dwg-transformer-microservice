import pika
from shapely.geometry import Point
import time

def callback(ch, method, properties, body):
    print(" [x] Received %r" % body)

def main():
    time.sleep(5)
    connection = pika.BlockingConnection(pika.ConnectionParameters('rabbitmq'))
    point = Point(0.0, 0.0)
    channel = connection.channel()
    channel.queue_declare(queue='CheckDWG', durable=True)
    channel.basic_publish(exchange='',
                          routing_key='CheckDWG',
                          body=point.wkt)
    print(" [x] Sent {}".format(point.wkt))

    channel.basic_consume(queue='CheckDWG',
                          auto_ack=True,
                          on_message_callback=callback)
    channel.start_consuming()

if __name__ == "__main__" :
    main()
