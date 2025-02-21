import React, { useState } from "react";
import { login } from "../api";
import { useNavigate } from "react-router-dom";

const Login: React.FC = () => {
  const [user, setUser] = useState({ email: "", password: "" });
  const navigate = useNavigate();

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    const response = await login(user);
    localStorage.setItem("token", response.data.token);
    navigate("/");
  };

  return (
    <div>
      <h2>Login</h2>
      <form onSubmit={handleSubmit}>
        <input type="email" placeholder="Email" onChange={(e) => setUser({ ...user, email: e.target.value })} required />
        <input type="password" placeholder="Password" onChange={(e) => setUser({ ...user, password: e.target.value })} required />
        <button type="submit">Login</button>
      </form>
    </div>
  );
};

export default Login;
