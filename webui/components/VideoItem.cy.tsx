import React from "react";
import VideoItem from "./VideoItem";
import { render, screen } from "@testing-library/react";
import { IVideo } from "../models/video";
const video = {
  id: 1,
  title: "title",
  description: "description",
  user: {
    email: "test@gmail.com",
    id: "ee456",
  },
  videoId: "ygkkN4Bxm38",
} as IVideo;

describe("<VideoItem />", () => {
  it("renders", () => {
    render(<VideoItem video={video} />);
  });

  it("should render a video title", () => {
    render(<VideoItem video={video} />);
    expect(screen.getByTestId('title')).to.text("title")
  });

  it("should render a video description", () => {
    render(<VideoItem video={video} />);
    expect(screen.getByTestId('description')).to.text("description")
  });

  it("should render a video share by with prefix text", () => {
    render(<VideoItem video={video} />);
    expect(screen.getByTestId('share-by')).to.text("Shared by test@gmail.com")
  });
});
