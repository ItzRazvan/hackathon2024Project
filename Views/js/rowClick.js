function rowClickMonth(event) {
	var target = event.target || event.srcElement;
	var id =
		target.parentElement.getElementsByClassName("tableIdLuna")[0].innerText;
	window.location.href = "/users?id=" + id;
}

function rowClickYear(event) {
	var target = event.target || event.srcElement;
	var id =
		target.parentElement.getElementsByClassName("tableIdAn")[0].innerText;
	window.location.href = "/users?id=" + id;
}
