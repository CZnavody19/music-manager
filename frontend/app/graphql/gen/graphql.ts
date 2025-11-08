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
  __typename?: 'DiscordConfig';
  webhookURL: Scalars['String']['output'];
};

export type DiscordConfigInput = {
  webhookURL: Scalars['String']['input'];
};

export type LoginInput = {
  password: Scalars['String']['input'];
  username: Scalars['String']['input'];
};

export type Mutation = {
  __typename?: 'Mutation';
  changeLogin: Scalars['Boolean']['output'];
  enableDiscord: Scalars['Boolean']['output'];
  enablePlex: Scalars['Boolean']['output'];
  enableYoutube: Scalars['Boolean']['output'];
  login: Scalars['String']['output'];
  logout: Scalars['Boolean']['output'];
  refreshPlexLibrary: Scalars['Boolean']['output'];
  sendTestDiscordMessage: Scalars['Boolean']['output'];
  setDiscordConfig: Scalars['Boolean']['output'];
  setPlexConfig: Scalars['Boolean']['output'];
  setYoutubeConfig: Scalars['Boolean']['output'];
  test: Scalars['Boolean']['output'];
};


export type MutationChangeLoginArgs = {
  input: LoginInput;
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


export type MutationLoginArgs = {
  input: LoginInput;
};


export type MutationSetDiscordConfigArgs = {
  config: DiscordConfigInput;
};


export type MutationSetPlexConfigArgs = {
  config: PlexConfigInput;
};


export type MutationSetYoutubeConfigArgs = {
  config: YoutubeConfigInput;
};

export type PlexConfig = {
  __typename?: 'PlexConfig';
  host: Scalars['String']['output'];
  libraryID: Scalars['Int']['output'];
  port: Scalars['Int']['output'];
  protocol: Scalars['String']['output'];
  token: Scalars['String']['output'];
};

export type PlexConfigInput = {
  host: Scalars['String']['input'];
  libraryID: Scalars['Int']['input'];
  port: Scalars['Int']['input'];
  protocol: Scalars['String']['input'];
  token: Scalars['String']['input'];
};

export type Query = {
  __typename?: 'Query';
  getDiscordConfig: DiscordConfig;
  getPlexConfig: PlexConfig;
  getServiceStatus: ServiceStatus;
  getVideosInPlaylist: Array<YouTubeVideo>;
  getYoutubeConfig: YoutubeConfig;
};

export type ServiceStatus = {
  __typename?: 'ServiceStatus';
  discord: Scalars['Boolean']['output'];
  plex: Scalars['Boolean']['output'];
  youtube: Scalars['Boolean']['output'];
};

export type YouTubeVideo = {
  __typename?: 'YouTubeVideo';
  channelTitle: Scalars['String']['output'];
  duration: Scalars['Int']['output'];
  id: Scalars['String']['output'];
  position: Scalars['Int']['output'];
  thumbnailUrl: Scalars['String']['output'];
  title: Scalars['String']['output'];
};

export type YoutubeConfig = {
  __typename?: 'YoutubeConfig';
  playlistID: Scalars['String']['output'];
};

export type YoutubeConfigInput = {
  playlistID: Scalars['String']['input'];
};
