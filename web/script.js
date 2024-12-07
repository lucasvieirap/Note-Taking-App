const NoteListElement = document.getElementById("note-list");
const ErrorMessageElement = document.getElementById("error-message");

async function getData(url) {
	try {
		const response = await fetch(url);
		if (!response.ok) {
			throw new Error(`Response Status: ${response.status}`);
		}

		const json = await response.json();
		return json;
	} catch (err) {
		console.log(err.message);
		ErrorMessageElement.innerText = `Error on server: ${err.message}`;
	}
}

async function main() {
	const NoteList = await getData("http://localhost:8080/list");
	NoteList.noteList.map(Note => {
		console.log(Note);
		NoteTitle = Note.note.split("=")[0] || Note.note.split("#")[0]
		NoteListElement.innerHTML += `<p>${NoteTitle}</p>`
	})
}

main();
