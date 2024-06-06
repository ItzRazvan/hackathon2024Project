fetch("/getData", {
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
				},
			});
		});
	}
});
