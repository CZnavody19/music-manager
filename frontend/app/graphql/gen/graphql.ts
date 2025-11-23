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
  Time: { input: string; output: string; }
  UUID: { input: any; output: any; }
};

export type DiscordConfig = {
  __typename?: 'DiscordConfig';
  webhookURL: Scalars['String']['output'];
};

export type DiscordConfigInput = {
  webhookURL: Scalars['String']['input'];
};

export type GeneralConfig = {
  __typename?: 'GeneralConfig';
  downloadPath: Scalars['String']['output'];
  tempPath: Scalars['String']['output'];
};

export type GeneralConfigInput = {
  downloadPath: Scalars['String']['input'];
  tempPath: Scalars['String']['input'];
};

export type LoginInput = {
  password: Scalars['String']['input'];
  username: Scalars['String']['input'];
};

export type Mutation = {
  __typename?: 'Mutation';
  changeLogin: Scalars['Boolean']['output'];
  deletePlexTrack: Scalars['Boolean']['output'];
  deleteTrack: Scalars['Boolean']['output'];
  download: Scalars['Boolean']['output'];
  enableDiscord: Scalars['Boolean']['output'];
  enablePlex: Scalars['Boolean']['output'];
  enableTidal: Scalars['Boolean']['output'];
  enableYoutube: Scalars['Boolean']['output'];
  login: Scalars['String']['output'];
  logout: Scalars['Boolean']['output'];
  matchVideo: Scalars['Boolean']['output'];
  refreshPlaylist: Scalars['Boolean']['output'];
  refreshPlexLibrary: Scalars['Boolean']['output'];
  refreshPlexTracks: Scalars['Boolean']['output'];
  sendTestDiscordMessage: Scalars['Boolean']['output'];
  setDiscordConfig: Scalars['Boolean']['output'];
  setGeneralConfig: Scalars['Boolean']['output'];
  setPlexConfig: Scalars['Boolean']['output'];
  setTidalConfig: Scalars['Boolean']['output'];
  setYoutubeConfig: Scalars['Boolean']['output'];
};


export type MutationChangeLoginArgs = {
  input: LoginInput;
};


export type MutationDeletePlexTrackArgs = {
  id: Scalars['ID']['input'];
};


export type MutationDeleteTrackArgs = {
  id: Scalars['UUID']['input'];
};


export type MutationEnableDiscordArgs = {
  enable: Scalars['Boolean']['input'];
};


export type MutationEnablePlexArgs = {
  enable: Scalars['Boolean']['input'];
};


export type MutationEnableTidalArgs = {
  enable: Scalars['Boolean']['input'];
};


export type MutationEnableYoutubeArgs = {
  enable: Scalars['Boolean']['input'];
};


export type MutationLoginArgs = {
  input: LoginInput;
};


export type MutationMatchVideoArgs = {
  trackID: Scalars['UUID']['input'];
  videoID: Scalars['String']['input'];
};


export type MutationSetDiscordConfigArgs = {
  config: DiscordConfigInput;
};


export type MutationSetGeneralConfigArgs = {
  config: GeneralConfigInput;
};


export type MutationSetPlexConfigArgs = {
  config: PlexConfigInput;
};


export type MutationSetTidalConfigArgs = {
  config: TidalConfigInput;
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

export type PlexTrack = {
  __typename?: 'PlexTrack';
  artist: Scalars['String']['output'];
  duration: Scalars['Int']['output'];
  id: Scalars['ID']['output'];
  mbid?: Maybe<Scalars['UUID']['output']>;
  title: Scalars['String']['output'];
  trackID?: Maybe<Scalars['UUID']['output']>;
};

export type Query = {
  __typename?: 'Query';
  getDiscordConfig: DiscordConfig;
  getGeneralConfig: GeneralConfig;
  getPlexConfig: PlexConfig;
  getPlexTracks: Array<PlexTrack>;
  getServiceStatus: ServiceStatus;
  getTidalConfig: TidalConfig;
  getTracks: Array<Track>;
  getVideoByID: YouTubeVideo;
  getVideosInPlaylist: Array<YouTubeVideo>;
  getYoutubeConfig: YoutubeConfig;
};


export type QueryGetVideoByIdArgs = {
  videoID: Scalars['String']['input'];
};

export type ServiceStatus = {
  __typename?: 'ServiceStatus';
  discord: Scalars['Boolean']['output'];
  plex: Scalars['Boolean']['output'];
  tidal: Scalars['Boolean']['output'];
  youtube: Scalars['Boolean']['output'];
};

export type Subscription = {
  __typename?: 'Subscription';
  tasks: Task;
};

export type Task = {
  __typename?: 'Task';
  ended: Scalars['Boolean']['output'];
  startedAt: Scalars['Time']['output'];
  title: Scalars['String']['output'];
};

export type TidalConfig = {
  __typename?: 'TidalConfig';
  audioQuality: Scalars['String']['output'];
  authAccessToken: Scalars['String']['output'];
  authClientID: Scalars['String']['output'];
  authClientSecret: Scalars['String']['output'];
  authExpiresAt: Scalars['Time']['output'];
  authRefreshToken: Scalars['String']['output'];
  authTokenType: Scalars['String']['output'];
  directoryPermissions: Scalars['Int']['output'];
  downloadRetries: Scalars['Int']['output'];
  downloadThreads: Scalars['Int']['output'];
  downloadTimeout: Scalars['Int']['output'];
  filePermissions: Scalars['Int']['output'];
  group: Scalars['Int']['output'];
  owner: Scalars['Int']['output'];
};

export type TidalConfigInput = {
  audioQuality: Scalars['String']['input'];
  authAccessToken: Scalars['String']['input'];
  authClientID: Scalars['String']['input'];
  authClientSecret: Scalars['String']['input'];
  authExpiresAt: Scalars['Time']['input'];
  authRefreshToken: Scalars['String']['input'];
  authTokenType: Scalars['String']['input'];
  directoryPermissions: Scalars['Int']['input'];
  downloadRetries: Scalars['Int']['input'];
  downloadThreads: Scalars['Int']['input'];
  downloadTimeout: Scalars['Int']['input'];
  filePermissions: Scalars['Int']['input'];
  group: Scalars['Int']['input'];
  owner: Scalars['Int']['input'];
};

export type Track = {
  __typename?: 'Track';
  artist: Scalars['String']['output'];
  id: Scalars['UUID']['output'];
  isrcs: Array<Scalars['String']['output']>;
  length: Scalars['Int']['output'];
  linkedPlex: Scalars['Boolean']['output'];
  linkedYoutube: Scalars['Boolean']['output'];
  title: Scalars['String']['output'];
};

export type YouTubeVideo = {
  __typename?: 'YouTubeVideo';
  channelTitle: Scalars['String']['output'];
  duration: Scalars['Int']['output'];
  id: Scalars['String']['output'];
  linked: Scalars['Boolean']['output'];
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
