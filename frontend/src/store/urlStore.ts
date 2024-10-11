import type { UrlShortenModel } from "@/models/UrlShorten.model";
import { atom } from "nanostores";

export const urlCreated = atom<UrlShortenModel>({ url: "", createdAt: "" });
