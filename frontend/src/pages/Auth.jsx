import LoginForm from "../components/LoginForm";
import SignupForm from "../components/SignUpForm";


const AuthPage = () => {
  return (
    <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100">
      <SignupForm />
      <div className="my-4"></div>
      <LoginForm />
    </div>
  );
};

export default AuthPage;
