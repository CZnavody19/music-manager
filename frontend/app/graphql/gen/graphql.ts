/* eslint-disable */
export type Maybe<T> = T | null;
export type InputMaybe<T> = T | null | undefined;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
};

export type DiscordConfig = {
  webhookURL: Scalars['String']['input'];
};

export type Mutation = {
  __typename?: 'Mutation';
  enableDiscord: Scalars['Boolean']['output'];
  enablePlex: Scalars['Boolean']['output'];
  enableYoutube: Scalars['Boolean']['output'];
  refreshPlexLibrary: Scalars['Boolean']['output'];
  sendTestDiscordMessage: Scalars['Boolean']['output'];
  setDiscordConfig: Scalars['Boolean']['output'];
  setPlexConfig: Scalars['Boolean']['output'];
};


export type MutationEnableDiscordArgs = {
  enable: Scalars['Boolean']['input'];
};


export type MutationEnablePlexArgs = {
  enable: Scalars['Boolean']['input'];
};


export type MutationEnableYoutubeArgs = {
  enable: Scalars['Boolean']['input'];
};


export type MutationSetDiscordConfigArgs = {
  config: DiscordConfig;
};


export type MutationSetPlexConfigArgs = {
  config: PlexConfig;
};

export type PlexConfig = {
  host: Scalars['String']['input'];
  libraryID: Scalars['Int']['input'];
  port: Scalars['Int']['input'];
  protocol: Scalars['String']['input'];
  token: Scalars['String']['input'];
};

export type Query = {
  __typename?: 'Query';
  getServiceStatus: ServiceStatus;
};

export type ServiceStatus = {
  __typename?: 'ServiceStatus';
  discord: Scalars['Boolean']['output'];
  plex: Scalars['Boolean']['output'];
  youtube: Scalars['Boolean']['output'];
};
