export function scrollToLastMessage() {
    const el = document.querySelector("#messages")

    if (el) {
        el.scrollTop = el.scrollHeight
    }
}
