# ProducerConsumerProblemGolang
## This repository manages backend for Hotel reservation.
#### *  Whenever user calls api endpoint to book a hotel room a go routine is called, which updates room state after 4 seconds, during these 4 seconds the shared resource(An array representing state of all rooms) is blocked. ####
#### *  Any request to any endpoint (checkIn or checkOut) is made to wait untill the previous one is complete. ####
#### *  Hence the producer(The Api caller) has to wait untill the consumer(the Api Endpoint) is ready to process the request. ####
