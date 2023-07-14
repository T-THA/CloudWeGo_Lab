namespace go student

//--------------------request & response--------------

struct Student {
    1: required i32 id(api.body='id'),
    2: required string name(api.body='name'),
    3: required string college(api.body='college'),
    4: optional list<string> email(api.body='email'),
}

struct RegisterResp {
    1: bool success(api.body='success'),
    2: string message(api.body='message'),
}

struct QueryReq {
    1: required i32 id(api.query='id')
}

//----------------------service-------------------
service StudentService {
    RegisterResp Register(1: Student student)(api.post = '/add-student-info')
    Student Query(1: QueryReq req)(api.get = '/query')
}