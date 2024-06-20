//make a function that numbers the absences like :#1, #2, #3, etc. It gets run when the page loads

window.onload = function () {
	var absences = document.getElementsByClassName("absence");
	for (var i = 0; i < absences.length; i++) {
		absences[i].innerHTML = "#" + (i + 1);
		console.log(absences[i]);
	}

	//get the id of the user from the url
	var url = new URL(window.location.href);
	var id = url.searchParams.get("id");

	fetch("/getData?id=" + id, {
		method: "GET",
		credentials: "include",
	}).then((response) => {
		if (response.ok) {
			response.json().then((data) => {
				var ctx = document.getElementById("myChart").getContext("2d");
				var myChart = new Chart(ctx, {
					type: "line",
					data: {
						labels: data.labels,
						datasets: [
							{
								label: "Absente",
								data: data.data,
								backgroundColor: ["rgba(255, 99, 132, 0.2)"],
								borderColor: ["rgba(255, 99, 132, 1)"],
							},
						],
					},
					options: {
						scales: {
							y: {
								beginAtZero: true,
							},
						},

						title: {
							display: true,
							text: "Numarul de absente in fiecare luna",
						},
						layout: {
							padding: {
								right: 30,
							},
						},

						responsive: true,
						maintainAspectRatio: false,
					},
				});
			});
		}
	});
};

document.getElementById("goBack").addEventListener("click", () => {
	window.location.href = "/tabele";
});
