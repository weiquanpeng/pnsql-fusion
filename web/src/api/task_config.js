import request from "@/request/index.js"

export const loadtaskconfig = ()=>{
    return request.get("/v1/task/list")
}

export const getSubTaskData = (id) => {
    return request.get(`/v1/task/subtask/list?id=${id}`);
};

export const loadOwnerTaskConfig = (owner)=>{
    return request.post("/v1/task/mine/list",{ owner: owner })
}

export const loadApproveTaskConfig = (owner)=>{
    return request.post("/v1/task/approve/list",{ owner: owner })
}

export const uptSubTaskConfigStatus = (id)=>{
    return request.post("/v1/task/approve/uptStatus",{ id: id })
}
