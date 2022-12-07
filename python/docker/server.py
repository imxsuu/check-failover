from http.server import BaseHTTPRequestHandler, HTTPServer
import logging
import sys
from datetime import datetime
import json
import threading

class S(BaseHTTPRequestHandler):

    def _set_response(self):
        self.send_response(200)
        self.send_header('Content-type', 'text/html')
        self.end_headers()

    def do_GET(self):
        logging.info("GET request,\nPath: %s\nHeaders:\n%s\n", str(self.path), str(self.headers))
        self._set_response()
        self.wfile.write("GET request for {}".format(self.path).encode('utf-8'))

    def do_POST(self):
        now = datetime.now().strftime('%Y-%m-%d %H:%M:%S')

        content_length = int(self.headers['Content-Length']) # <--- Gets the size of data
        post_data = self.rfile.read(content_length) # <--- Gets the data itself
        data = json.loads(post_data)
        logging.info("[RECV] [" + now  + "]\n" + str(data) + "\n")        

        if data['deviceName'] == 'RaspberryPi+':
            def kill_me_please(server):
                server.shutdown()
            t = threading.Thread(kill_me_please, (httpd,))
            t.start()
            self.send_error(500)
        
def run(server_class=HTTPServer, handler_class=S, port=9010):
    logging.basicConfig(level=logging.INFO)
    server_address = ('', port)
    httpd = server_class(server_address, handler_class)
    logging.info('Starting httpd...\n')
    try:
        httpd.serve_forever()
    except:
        pass

    httpd.server_close()
    logging.info('Stopping httpd...\n')

if __name__ == '__main__':
    from sys import argv

    if len(argv) == 2:
        run(port=int(argv[1]))
    else:
        run()
