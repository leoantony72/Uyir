import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
// import LoginPage from "./pages/LoginPage";
// import TeacherDashboard from "./pages/TeacherDashboard.jsx";
import { useState } from "react";
// import { onAuthStateChanged } from "firebase/auth";
// import { auth } from "./firebase"; // Your Firebase config
import "./App.css";
import { SignUp } from "./pages/SignUp.jsx";
import { Login } from "./pages/Login.jsx";
import { ReportsPage } from "./pages/ReportsPage.jsx";


const App = () => {
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);

  // useEffect(() => {
  //   // Listen for auth state changes
  //   // const unsubscribe = onAuthStateChanged(auth, (currentUser) => {
  //   //   setUser(currentUser);
  //   //   setLoading(false); // Stop loading once the auth state is determined
  //   // });

  //   // Cleanup listener on unmount
  //   // return () => unsubscribe();
  // }, []);

  // if (loading) {
  //   return<div className="Loader">
  //   <div>Loading...</div><div className="loading-spinner"></div>
  // </div>;
  // }

  return (
    <Router>
      <Routes>
        <Route path="/signup" element={<SignUp />} />
        <Route path="/login" element={<Login />} />
        <Route path="/user" element={<ReportsPage />} />
        {/* <Route
          path="/faculty"
          element={
            <ProtectedRoute requiredRole="faculty">
              <TeacherDashboard />
            </ProtectedRoute>
          }
        /> */}
      </Routes>
    </Router>
  );
};

export default App;
