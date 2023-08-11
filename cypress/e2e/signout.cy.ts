/// <reference types="cypress" />

describe("Register Account", async () => {
  beforeEach(() => {
    cy.visit("http://localhost:3000/signup");
  });

  it("displays input all input by default", () => {
    cy.get("[name=email]").should("have.length", 1);
    cy.get("[name=password]").should("have.length", 1);
    cy.get("[name=confirmPassword]").should("have.length", 1);
  });
  it("Register success", () => {
    cy.get("[name=email]").type("bbbbbbb@g.com");
    cy.get("[name=password]").type("bbbbbbb@g.com");
    cy.get("[name=confirmPassword]").type("bbbbbbb@g.com");
    cy.get("button").click();
  });

  it("Register failed", () => {
    cy.get("[name=email]").type("bbbbbbbdd@g.com");
    cy.get("[name=password]").type("bbbbbbb@g.com");
    cy.get("[name=confirmPassword]").type("bbbbbvbb@g.com");
    cy.get("button").click();
  });

  it("Should API request return success", () => {
    cy.request("POST", "http://localhost:8090/v1/auth/register", {
      email: "thanhduongbkdn2012@gmail.com",
      password: "abcd1234",
    });
  });
});
