const PVA = {
    confirm(title,data,type,message){
        ElMessage.closeAll()
        return  ElMessageBox.confirm(
            data,
            title,
            {
                confirmButtonText: '确认',
                cancelButtonText: '取消',
                type: type,
            },
        ).then(() => {
                ElMessage({
                    type: 'success',
                    message: message,
                })
            })
    },

    notice(title,data,type){
        ElNotification.closeAll()
        return  ElNotification({
            title: title,
            message: data,
            type: type,
            dangerouslyUseHTMLString: true,
        })
    }
}

export default PVA