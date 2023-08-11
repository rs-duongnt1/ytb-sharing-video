import React, { useState } from "react";
import SignInForm from "./SignInForm";
import { render, screen } from "@testing-library/react";
import { BrowserRouter } from "react-router-dom";

describe("<SignInForm />", () => {
  it("renders", () => {
    const handleLogin = () => {};
    const loading = true;
    render(
      <BrowserRouter>
        <SignInForm handleLogin={handleLogin} loading={loading} />
      </BrowserRouter>
    );
  });

  it("should render all element", () => {
    const handleLogin = () => {};
    const loading = false;
    render(
      <BrowserRouter>
        <SignInForm handleLogin={handleLogin} loading={loading} />
      </BrowserRouter>
    );

    expect(screen.getByTestId('email')).to.exist;
    expect(screen.getByTestId('password')).to.exist;
    expect(screen.getByTestId('button')).to.exist;
  })

  it("should button is disabled", () => {
    const handleLogin = () => {};
    const loading = true;
    render(
      <BrowserRouter>
        <SignInForm handleLogin={handleLogin} loading={loading} />
      </BrowserRouter>
    );

    expect(screen.getByRole('button')).to.disabled;
  });
});
