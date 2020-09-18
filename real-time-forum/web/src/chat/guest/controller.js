export default class GuestController {
    constructor(model, view) {
        this.model = model
        this.view = view

        this.displayGuest()
    }

    displayGuest = () => {
        this.view.display(this.model.get())
    }

    delete = () => {
        this.view.delete()
    }
}