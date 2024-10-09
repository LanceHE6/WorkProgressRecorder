export const mapping = (item, data) => {
    if(item.isDate){
        return toLocaleString(data)
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

export function mobileDeviceViewport(){
    const meta = document.querySelector('meta[name="viewport"]');
    if (meta) {
        meta.content = 'width=device-width, initial-scale=1.0';
    } else {
        const newMeta = document.createElement('meta');
        newMeta.name = 'viewport';
        newMeta.content = 'width=device-width, initial-scale=1.0';
        document.head.appendChild(newMeta);
    }
}

export function screenDeviceViewport(){
    const viewportMeta = document.querySelector('meta[name="viewport"]');
    // 如果找到了，就移除它
    if (viewportMeta) {
        viewportMeta.parentNode.removeChild(viewportMeta);
    }
}

export function toLocaleString(date){
    if(!date){
        return null
    }
    return new Date(date).toLocaleString()
}

export function getReverseList(list){
    if(!list){
        return null
    }
    return [...list].reverse()
}