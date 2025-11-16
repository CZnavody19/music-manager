from internal.mq import NewMQ
from internal.compare import NewComparator
from internal.download import NewDownloader

def main():
    print("Starting Tidal downloader...")

    comparator = NewComparator()

    downloader = NewDownloader(comparator=comparator)

    mq = NewMQ(downloader=downloader)

    print("Tidal downloader started.")
    mq.run()
    print("Tidal downloader stopped.")

if __name__ == "__main__":
    main()