### Dashboard Service

#### Developer Notes:
Follow the installation steps:
- Make sure you have the `node` version `> 5.50`
  `$ node -v  # prints the current version`

- cd into the dashboard directory and run
  1. Run `$ npm install` to install all the app dependencies.
  2. Run `$ npm run start` to run the web application.
- If you are running the SAL user and app service on different ports
then you may have to update the IP and ports in `services/dashboard/app/services/index.js` constants.

- Access the application on http://localhost:8000
