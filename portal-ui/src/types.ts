// This file is part of MinIO Console Server
// Copyright (c) 2019 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

export interface SystemState {
  loggedIn: boolean;
  operatorMode: boolean;
  sidebarOpen: boolean;
  session: string;
  userName: string;
  serverNeedsRestart: boolean;
  serverIsLoading: boolean;
}

export const USER_LOGGED = "USER_LOGGED";
export const OPERATOR_MODE = "OPERATOR_MODE";
export const MENU_OPEN = "MENU_OPEN";
export const SERVER_NEEDS_RESTART = "SERVER_NEEDS_RESTART";
export const SERVER_IS_LOADING = "SERVER_IS_LOADING";

interface UserLoggedAction {
  type: typeof USER_LOGGED;
  logged: boolean;
}

interface OperatorModeAction {
  type: typeof OPERATOR_MODE;
  operatorMode: boolean;
}

interface SetMenuOpenAction {
  type: typeof MENU_OPEN;
  open: boolean;
}

interface ServerNeedsRestartAction {
  type: typeof SERVER_NEEDS_RESTART;
  needsRestart: boolean;
}

interface ServerIsLoading {
  type: typeof SERVER_IS_LOADING;
  isLoading: boolean;
}

export type SystemActionTypes =
  | UserLoggedAction
  | OperatorModeAction
  | SetMenuOpenAction
  | ServerNeedsRestartAction
  | ServerIsLoading;
