export const mapping = (item, data) => {
    if(item.isDate){
        return new Date(data).toLocaleString()
    }
    else if(item.isMapping){
        for(const i of item.mappingList){
            const item2 = i
            // 映射
            if(data === item2.value){
                return item2.label
            }
        }
    }
    else{
        return data
    }
}