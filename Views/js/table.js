document.getElementById("tableType").addEventListener("change", () => {
	var tableType = document.getElementById("tableType").value;
	if (tableType == "month") {
		document.getElementById("month").style.display = "table";
		document.getElementById("year").style.display = "none";
	}
	if (tableType == "year") {
		document.getElementById("month").style.display = "none";
		document.getElementById("year").style.display = "table";
	}
});

document.getElementById("goBack").addEventListener("click", () => {
	window.location.href = "/";
});
