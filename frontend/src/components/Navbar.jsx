const Navbar = ({ isLoggedIn, onLogout }) => {
  const handleLoginRedirect = () => {
    window.location.href = "/auth/login";
  };

  return (
    <nav className="bg-blue-600 p-4 flex justify-between items-center text-white">
      <h1 className="text-xl font-bold">My App</h1>
      <div>
        {isLoggedIn ? (
          <button onClick={onLogout} className="bg-red-500 px-4 py-2 rounded">Logout</button>
        ) : (
          <button onClick={handleLoginRedirect} className="bg-green-500 px-4 py-2 rounded">Login</button>
        )}
      </div>
    </nav>
  );
};

export default Navbar;
