import logo from "./logo.svg";
import "./App.css";
import { useEffect, useState } from "react";
import axios, { post } from "axios";

function App() {
  // useEffect(() => {
  //   fetch("http://localhost:3001/user/2")
  //     .then((res) => res.json())
  //     .then((json) => {
  //       console.log(json);
  //     });
  // }, []);

  const [file, setFile] = useState();

  const onFormSubmit = (e) => {
    e.preventDefault(); // Stop form submit
    fileUpload(file).then((response) => {
      console.log(response.data);
    });
  };

  const fileUpload = (file) => {
    const url = "http://localhost:8080/upload";
    console.log(file);
    const formData = new FormData();
    formData.append("file", file);
    formData.append("name", "gfdf");
    formData.append("email", "gdf@hgfd.hgfj");
    const config = {
      headers: {
        "content-type": "multipart/form-data",
      },
    };
    return post(url, formData, config);
  };

  const onChange = (e) => {
    setFile(e.target.files[0]);
  };

  return (
    <div>
      <h1>Bind file with fields</h1>
      <form>
        Name: <input type='text' name='name' />
        <br />
        Email: <input type='email' name='email' />
        <br />
        File: <input type='file' onChange={onChange} />
        <br />
        <br />
        <button type='submit' onClick={onFormSubmit}>
          Upload
        </button>
      </form>
    </div>
  );
}

export default App;
