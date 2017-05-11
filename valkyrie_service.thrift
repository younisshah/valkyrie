namespace go vservice

service ValkyrieService {
    bool send(1:string message, 2:string queue)
}