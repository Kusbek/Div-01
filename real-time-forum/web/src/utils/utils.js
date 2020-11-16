

export const displayModal = (msg) => {
    alert(msg)
}


export const createElement = (tag, ...classNames) => {
    const element = document.createElement(tag)
    if (classNames.length != 0) {
        classNames.forEach((className) => {
            element.classList.add(className)
        })
    }
    return element
}

export const partition = (arr, func) => {
    let truthy = []
    let falsey = []
    for(let i = 0; i<arr.length;i++){
        if (func(arr[i], i, arr)){
            truthy.push(arr[i])
        }else{
            falsey.push(arr[i])
        }
    }
    return [truthy,  falsey]
}

export const formatDate = (timestamp) => {
    let date = new Date(timestamp)
    var month = new Array();
    month[0] = "January";
    month[1] = "February";
    month[2] = "March";
    month[3] = "April";
    month[4] = "May";
    month[5] = "June";
    month[6] = "July";
    month[7] = "August";
    month[8] = "September";
    month[9] = "October";
    month[10] = "November";
    month[11] = "December";
    let seconds = date.getSeconds()>9?date.getSeconds():`0${date.getSeconds()}`
    return `${date.getDate()} ${month[date.getMonth()]} ${date.getHours()}:${date.getMinutes()}:${seconds}`
}



export const validateNickname = (nickname) => {
    if (nickname.length < 5 || nickname.length > 19) {
        return false
    }
    let re = new RegExp(/^[a-zA-Z0-9]+$/, "g")
    if (!nickname.match(re)) {
        return false
    }
    return true
}


export const validateEmail = (email) => {
    let re = new RegExp(/^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/, "g")
    if (re.test(email)) {
        return true
    }
    return false
}


export const validateName = (name) => {
    if (name.length < 1 || name.length > 20) {
        return false
    }
    let re = new RegExp(/^[a-zA-Z]+$/, "g")
    if (!re.test(name)) {
        return false
    }
    return true
}