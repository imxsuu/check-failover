from http.server import BaseHTTPRequestHandler, HTTPServer
from kubernetes import client, config
import logging
import sys

def create_pv_object(msg):
    spec=client.V1PersistentVolumeSpec(
        storage_class_name="manual",
        capacity={
            "storage": "1Gi"
        },
        volume_mode="Filesystem",
        access_modes=["ReadWriteOnce"],
        persistent_volume_reclaim_policy="Retain",
        local=client.V1LocalVolumeSource(
            path=f"/dat/gedge-pv"
        )
    )
    
    pv = client.V1PersistentVolume(
        api_version="v1",
        kind="PersistentVolume",
        metadata=client.V1ObjectMeta(
            name=f"gedge-pv"
        ),
        spec=spec,
    )

    return pv

def create_pv(api, pv):
    resp = api.create_persistent_volume(
        body=pv
    )

    logging.info("\n[INFO] Persistent Volume created.\n")
    logging.info("%s" % ("NAME"))
    logging.info(
        "%s\n"
        % (
            resp.metadata.name,
        )
    )
    

def create_pvc_object(msg):
    spec=client.V1PersistentVolumeClaimSpec(
        volume_name=f"gedge-pv",
        storage_class_name="manual",
        access_modes=["ReadWriteOnce"],
        resources=client.V1ResourceRequirements(
            requests={"storage": "10Gi"},
        ),
    )
    
    pvc = client.V1PersistentVolumeClaim(
        api_version="v1",
        kind="PersistentVolumeClaim",
        metadata=client.V1ObjectMeta(
            name=f"gedge-pv-claim",
            namespace="gedge-platfm"
        ),
        spec=spec,
    )

    return pvc

def create_pvc(api, pvc):
    resp = api.create_namespaced_persistent_volume_claim(
        body=pvc, namespace="agc2021"
    )

    logging.info("\n[INFO] Persistent Volume Claim created.\n")
    logging.info("%s" % ("NAME"))
    logging.info(
        "%s\n"
        % (
            resp.metadata.name,
        )
    )


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
        content_length = int(self.headers['Content-Length']) # <--- Gets the size of data
        post_data = self.rfile.read(content_length) # <--- Gets the data itself
        
        if post_data.decode('utf-8') == 'a' :

            logging.info("POST request,\nPath: %s\nHeaders:\n%s\n\nBody:\n%s\n",
                str(self.path), str(self.headers), post_data.decode('utf-8'))

            self._set_response()
            self.wfile.write("POST request for {}".format(self.path).encode('utf-8'))

        elif post_data.decode('utf-8') == 'b' :
            # apps_v1 = client.CoreV1Api()
            
            # pv = create_pv_object(post_data.decode('utf-8'))
            # create_pv(apps_v1, pv)

            # pvc = create_pvc_object(post_data.decode('utf-8'))
            # create_pvc(apps_v1, pvc)

            logging.fatal('hi')
            # try:
            #     raise Exception
            # except: 
            #     sys.exit()


            logging.info("POST request,\nPath: %s\nHeaders:\n%s\n\nBody:\n%s\n",
                str(self.path), str(self.headers), post_data.decode('utf-8'))
            logging.info("error")

            self._set_response()
            self.wfile.write("POST request for {}".format(self.path).encode('utf-8'))
        
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