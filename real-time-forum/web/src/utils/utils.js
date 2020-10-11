

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