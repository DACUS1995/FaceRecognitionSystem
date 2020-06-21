# FaceRecognitionSystem

Small and extensible face recognition system that can store face embeddings and find matches in received images.
Both the face recognition and the image sampler are separated and can be runned from separated machines.

---

### Components

#### Core
The central server has a json config file that can the edited according to the new setup.
You can use your prefered storage system by implementing the `DatabaseClient` interface from `core/dbactions/database.go`

#### Services
One service that handles the face detection and the embedding operation.
One service that collects images from a local available camera.

---

In the root of the project there are `.service` files to register every components as a systemd service. You just need to replace the `<ROOT PATH>` with project location in the filesystem. 
