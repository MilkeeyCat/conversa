export function toggleCreateRoomForm() {
    const form = document.querySelector("#createRoomForm")
    const button = document.querySelector("#createRoomButton")

    if (form && button) {
        form.classList.toggle("!block")
        button.classList.toggle("hidden")
    }
}

export function toggleJoinRoomForm() {
    const form = document.querySelector("#joinRoomForm")
    const button = document.querySelector("#joinRoomButton")

    if (form && button) {
        form.classList.toggle("!block")
        button.classList.toggle("hidden")
    }
}
