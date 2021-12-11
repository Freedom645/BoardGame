
import { FirebaseUIModule, firebase, firebaseui } from 'firebaseui-angular';

export const FirebaseUiAuthConfig: firebaseui.auth.Config = {
  autoUpgradeAnonymousUsers: false,
  signInFlow: 'redirect',
  signInOptions: [
    // firebase.auth.GoogleAuthProvider.PROVIDER_ID,
    // firebase.auth.TwitterAuthProvider.PROVIDER_ID,
    firebase.auth.EmailAuthProvider.PROVIDER_ID,
    firebaseui.auth.AnonymousAuthProvider.PROVIDER_ID
  ],
  // tosUrl: 'http://localhost:4200/room',
  // privacyPolicyUrl: 'プライバシーポリシーのURL',
  credentialHelper: firebaseui.auth.CredentialHelper.GOOGLE_YOLO,
};
