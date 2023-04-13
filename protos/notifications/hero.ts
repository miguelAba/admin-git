/* eslint-disable */
import { Observable } from "rxjs";

export const protobufPackage = "hero";

export enum Corpus {
  CORPUS_UNSPECIFIED = 0,
  CORPUS_UNIVERSAL = 1,
  CORPUS_WEB = 2,
  UNRECOGNIZED = -1,
}

export interface HeroById {
  id: number;
}

export interface SampleMessage {
  name?: string | undefined;
  subMessage?: Hero | undefined;
}

export interface Hero {
  id: number;
  name: string;
  corpus: Corpus;
  results: Corpus[];
  details: string[];
  projects: { [key: string]: Hero };
}

export interface Hero_ProjectsEntry {
  key: string;
  value: Hero | undefined;
}

export interface HeroesService {
  Unary(request: HeroById): Promise<Hero>;
  ServerStream(request: HeroById): Observable<Hero>;
  ClientStream(request: Observable<HeroById>): Promise<Hero>;
  ClientAndServerStream(request: Observable<HeroById>): Observable<Hero>;
}
