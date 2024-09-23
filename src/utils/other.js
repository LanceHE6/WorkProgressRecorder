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
    else if(item.isMappingList){
        let result = ''
        for(const i in data){
            for(const j of item.mappingList){
                const item2 = j
                // 映射
                if(data[i] === item2.value){
                   result += `${item2.label}&`
                }
            }
        }
        return result.slice(0, -1)
    }
    else{
        return data
    }
}