namespace go hello.example

struct HelloReq {
    1: string Name (api.query="name");
}

struct HelloResp {
    1: string RespBody;
}

struct AddReq{
    1: i32 Id (api.body="Id");
    2: string Name (api.body="Name");
    3: string Favorite (api.body="Favorite");
}

struct AddResp {
    1: string RespBody;
}

struct StudentReq{
    1: i32 Id (api.query="id");
    
}
struct StudentResp{
    1: string RespBody;
}

service HelloService {
    HelloResp HelloMethod(1: HelloReq request) (api.get="/hello");
}

service NewService {
    HelloResp NewMethod(1: HelloReq request) (api.get="/new");
}

service StudentService{
    StudentResp StudentMethod(1: StudentReq request) (api.get="/query");
    AddResp StudentAddMethod(1: AddReq request) (api.post="/add");
}