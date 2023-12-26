export function toggleCreateRoomForm() {
    const form = document.querySelector("#createRoomForm")
    const button = document.querySelector("#createRoomButton")

    if (form && button) {
        form.classList.toggle("!block")
        button.classList.toggle("hidden")
    }
}
