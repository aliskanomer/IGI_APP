import ReactDOM from "react-dom/client";
import App from "./App";

const root = ReactDOM.createRoot(
  document.getElementById("root") as HTMLElement
);

//*1
root.render(<App />);

// *1 There should be react.Strict wrapping the App for dev env but it renders everything twice and makes every call twice and it makes everything harder to debug
