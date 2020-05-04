# FaceRecognitionSystem

Small and extensible face recognition system that can store face embeddings and find matches in received images.
Both the face recognition and the image sampler are separated and can be runned from separated machines.

---

### Setup

#### Core
The central server has a json config file that can the edited according to the new setup.
You can use your prefered storage system by implementing the `DatabaseClient` interface from `core/dbactions/database.go`

#### Services
TODO Add info
