## Hackathon: in4it 2024 Project

Me and my friend participated in the 2024 edition of the in4it hackathon (Romania).

This is half of the project, which consisted of a face detection app in python(he did it), and an online catalogue which communicates with the face detection api through requests.
In the catalogue, you have two type of accounts: admin and user

*Admin:* Can add users - you have to upload a photo too, and the photo will be transmitted through a request to the api. I will also send the user id that i created. When the user is absent, the api will tell me his id, and I my app will mark him as absent
      
*Admin:* Can see and search through a table of all the users. Then, the admin can click a user, and see all his absences in order.
      
*User:* Can see his absences + a graph of them

---------------------------------------------------------------------

This was created for schools to use. If a students are absent for more than 50% of the class, the api will send me a request of them, and they will be automatically be marked as absent.
This makes it easier for the teacher, and safer for the stundents, who can't run from school anymore without being caught.

The api can also be used by workspaces. There will be a camera at the entrance of the building. If someone doesn't show up to work, the api will send the request and the employes will be marked
