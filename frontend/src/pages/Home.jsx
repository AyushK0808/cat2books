const Home = ({ isLoggedIn }) => {
  return (
    <div className="flex items-center justify-center h-screen text-3xl font-semibold">
      {isLoggedIn ? "Yoo" : "Hello"}
    </div>
  );
};

export default Home;
