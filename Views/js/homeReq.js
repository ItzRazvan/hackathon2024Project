//make the singup button visible only if the user is admin

fetch("/user", {
	method: "GET",
	credentials: "include",
}).then((response) => {
	if (response.ok) {
		response.json().then((user) => {
			if (user.access == "admin") {
				document.getElementById("goToSignup").style.display = "block";
				document.getElementById("goToTable").style.display = "block";
			}
		});
	}
});

document.getElementById("goToLogin").addEventListener("click", () => {
	fetch("/login", {
		method: "DELETE",
		credentials: "include",
	}).then((response) => {
		if (response.ok) {
			window.location.href = "/login";
		}
	});
});

document.getElementById("goToSignup").addEventListener("click", () => {
	window.location.href = "/signup";
});

document.getElementById("goToTable").addEventListener("click", () => {
	window.location.href = "/tabele";
});
