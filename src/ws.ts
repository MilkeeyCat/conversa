type HTMXMessage = {
    message: string;
    id: number;
    //Some other fields
}

let currentRoomId = -1;

export function setCurrentRoomId(id: number) {
    currentRoomId = id;
    console.log(currentRoomId)
}

window.addEventListener("htmx:wsConfigSend", (e: Event) => {
    //@ts-ignore
    e.detail.parameters.id = currentRoomId;
    //@ts-ignore
    console.log(e.detail.parameters)
})
