document.getElementById("signupForm").addEventListener("submit", () => {
	event.preventDefault();
	document.getElementById("signupError").innerText = "";
	document.getElementById("signupSuccess").innerText = "";

	fetch("/signup", {
		method: "POST",
		credentials: "include",
		body: new FormData(document.getElementById("signupForm")),
	})
		.then((response) => {
			if (response.ok) {
				document.getElementById("name").value = "";
				document.getElementById("email").value = "";
				document.getElementById("password").value = "";
				document.getElementById("postImg").value = "";

				document.getElementById("signupSuccess").innerText =
					"Contul se creeaza!";
			} else if (response.status == 403) {
				document.getElementById("signupError").innerText =
					"Email deja folosit sau parola prea slaba";
			} else if (response.status == 400) {
				document.getElementById("signupError").innerText =
					"Eroare la conectarea cu serverul";
			} else {
				document.getElementById("signupError").innerText = "Eroare";
			}
		})
		.catch((error) => {
			console.error("Error:", error);
		});
});

document.getElementById("goBack").addEventListener("click", () => {
	window.location.href = "/";
});
