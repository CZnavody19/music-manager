from pika.adapters.blocking_connection import BlockingChannel
from pika import BlockingConnection, URLParameters
from pika.spec import Basic, BasicProperties

from utils import get_env_var
from .download import Downloader, NewDownloader
from json import loads, dumps
from models.track import Track

class MQ:
    init_chan: BlockingChannel
    connection: BlockingConnection
    downloader: Downloader

    def __init__(self, connection: BlockingConnection, downloader: Downloader, init_chan: BlockingChannel):
        self.connection = connection
        self.downloader = downloader
        self.init_chan = init_chan

    def download_request(self, ch: BlockingChannel, method: Basic.Deliver, properties: BasicProperties, body: bytes):
        job = loads(body)
        tr = Track(job)

        ch.basic_ack(delivery_tag=method.delivery_tag)
        chan = self.connection.channel()

        try:
            path = self.downloader.download(tr)

            assert path is not None

            print(f"Download complete for track ID {tr.ID}, saved to {path.absolute().as_posix()}")

            chan.basic_publish("downloads_complete", "success", dumps({
                "track": tr.json(),
                "file_path": path.absolute().as_posix(),
                "error": None
            }))

            return

        except Exception as e:
            print(f"Download failed for track ID {tr.ID}")
            chan.basic_publish("downloads_complete", "fail", dumps({
                "track": tr.json(),
                "file_path": None,
                "error": "Download failed"
            }))

    def reload_req(self, ch: BlockingChannel, method: Basic.Deliver, properties: BasicProperties, body: bytes):
        ch.basic_ack(delivery_tag=method.delivery_tag)

        print("Reloading downloader configuration")

        self.downloader = NewDownloader(self.downloader.comparator) # Refetches config

    def run(self):
        self.init_chan.start_consuming()

def NewMQ(downloader: Downloader) -> MQ:
    connection = BlockingConnection(URLParameters(get_env_var("RABBITMQ_URL")))

    chan = connection.channel()

    chan.exchange_declare(exchange='downloads', exchange_type='direct', durable=True)
    chan.exchange_declare(exchange='downloads_complete', exchange_type='direct', durable=True)
    chan.exchange_declare(exchange='reload', exchange_type='direct', durable=True)

    chan.queue_declare(queue='download.tidal', durable=True)
    chan.queue_declare(queue='reload.tidal', durable=True)

    chan.queue_bind(exchange='downloads', queue='download.tidal', routing_key='tidal')
    chan.queue_bind(exchange='reload', queue='reload.tidal', routing_key='tidal')

    mq = MQ(
        init_chan=chan,
        connection=connection,
        downloader=downloader
    )

    chan.basic_consume(queue='download.tidal', on_message_callback=mq.download_request)
    chan.basic_consume(queue='reload.tidal', on_message_callback=mq.reload_req)

    return mq