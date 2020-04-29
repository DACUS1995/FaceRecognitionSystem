from server import Server
from config import Config

def main():
	service_server = Server(
		address = Config.server_address,
		port = Config.server_port
	)
	service_server.run()


if __name__  == "__main__":
	main()