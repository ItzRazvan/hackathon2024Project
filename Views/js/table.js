document.getElementById("tableType").addEventListener("change", () => {
	var tableType = document.getElementById("tableType").value;
	if (tableType == "month") {
		document.getElementById("month").style.display = "table";
		document.getElementById("year").style.display = "none";
		document.getElementById("search").value = "";
		var rows = document.getElementsByClassName("coloanaLuna");
		for (var i = 0; i < rows.length; i++) {
			rows[i].style.display = "";
		}
	}
	if (tableType == "year") {
		document.getElementById("month").style.display = "none";
		document.getElementById("year").style.display = "table";
		document.getElementById("search").value = "";
		var rows = document.getElementsByClassName("coloanaAn");
		for (var i = 0; i < rows.length; i++) {
			rows[i].style.display = "";
		}
	}
});

document.getElementById("goBack").addEventListener("click", () => {
	window.location.href = "/";
});

function searchTable() {
	var input, filter, tableType, numeLuna, numeAn, i, txtValue;
	input = document.getElementById("search");
	filter = input.value.toUpperCase();
	tableType = document.getElementById("tableType").value;
	numeLuna = document.getElementsByClassName("numeLuna");
	numeAn = document.getElementsByClassName("numeAn");
	if (tableType === "month") {
		for (i = 0; i < numeLuna.length; i++) {
			txtValue = numeLuna[i].textContent || numeLuna[i].innerText;
			if (txtValue.toUpperCase().indexOf(filter) > -1) {
				numeLuna[i].parentNode.style.display = "";
			} else {
				numeLuna[i].parentNode.style.display = "none";
			}
		}
	} else if (tableType === "year") {
		for (i = 0; i < numeAn.length; i++) {
			txtValue = numeAn[i].textContent || numeAn[i].innerText;
			if (txtValue.toUpperCase().indexOf(filter) > -1) {
				numeAn[i].parentNode.style.display = "";
			} else {
				numeAn[i].parentNode.style.display = "none";
			}
		}
	}
}
