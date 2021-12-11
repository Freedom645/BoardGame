// This file can be replaced during build by using the `fileReplacements` array.
// `ng build` replaces `environment.ts` with `environment.prod.ts`.
// The list of file replacements can be found in `angular.json`.

export const environment = {
  production: false,
  siteName: "local",
  api: {
    http: "http://",
    ws: "ws://",
    host: "192.168.1.52:8080",
  },
  firebaseConfig: {
    apiKey: "AIzaSyC3CW9b6a6iUmr2TmV10hbcLPjBOWtyAPk",
    authDomain: "board-game-a5752.firebaseapp.com",
    projectId: "board-game-a5752",
    storageBucket: "board-game-a5752.appspot.com",
    messagingSenderId: "216936644273",
    appId: "1:216936644273:web:cf67216c5596948d8c2c90",
    measurementId: "${config.measurementId}"
  },
};

/*
 * For easier debugging in development mode, you can import the following file
 * to ignore zone related error stack frames such as `zone.run`, `zoneDelegate.invokeTask`.
 *
 * This import should be commented out in production mode because it will have a negative impact
 * on performance if an error is thrown.
 */
// import 'zone.js/plugins/zone-error';  // Included with Angular CLI.
