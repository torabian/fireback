import { AxiosInstance, AxiosRequestConfig, AxiosResponse } from 'axios';
import { ExecutionContext, createParamDecorator } from '@nestjs/common';
import { TypedRequestInit, URLSearchParamsX, fetchx } from './sdk';
/**
* Action to communicate with the action classicSignup
*/


	/**
 * FetchClassicSignupAction
 */

export class FetchClassicSignupAction {
  static URL = '/passports/signup/classic';
  static Method = 'post';

  
  	static Axios = (clientInstance: AxiosInstance, config: AxiosRequestConfig<unknown>) =>
		clientInstance
		.request<unknown, AxiosResponse<unknown>, unknown>(config)

		
		.then((res) => {
			return {
			...res,

			
			// if there is a output class, create instance out of it.
			data: new ClassicSignupRes(res.data),
			};
		});
		
	
  
  
	static Fetch = (
		init?: TypedRequestInit<unknown, unknown> | undefined,
		qs?: ClassicSignupQueryParams,
		overrideUrl?: string
	) =>
		fetchx<unknown, unknown, unknown>(
			new URL((overrideUrl ?? FetchClassicSignupAction.URL ) + '?' + qs?.toString()),
			init
		)

		
			.then((res) => res.json())
		

	
		
			.then((data) => new ClassicSignupRes (data));
		

	
}









	
/**
  * @decription The base class definition for classicSignupReq
  **/

export class ClassicSignupReq {
	 
		/**
  * @type {string}
  * @description 
  **/
 value;
		/**
  * @returns {string}
  * @description 
  **/
getValue () { return this[`value`] }
		/**
  * @param {string}
  * @description 
  **/
setValue (value) { this[`value`] = value; return this; } 
	 
		/**
  * @type {string}
  * @description Required when the account creation requires recaptcha, or otp approval first. If such requirements are there, you first need to follow the otp apis, get the session secret and pass it here to complete the setup.
  **/
 sessionSecret;
		/**
  * @returns {string}
  * @description Required when the account creation requires recaptcha, or otp approval first. If such requirements are there, you first need to follow the otp apis, get the session secret and pass it here to complete the setup.
  **/
getSessionSecret () { return this[`sessionSecret`] }
		/**
  * @param {string}
  * @description Required when the account creation requires recaptcha, or otp approval first. If such requirements are there, you first need to follow the otp apis, get the session secret and pass it here to complete the setup.
  **/
setSessionSecret (value) { this[`sessionSecret`] = value; return this; } 
	 
		/**
  * @type {"phonenumber" | "email"}
  * @description 
  **/
 type;
		/**
  * @returns {"phonenumber" | "email"}
  * @description 
  **/
getType () { return this[`type`] }
		/**
  * @param {"phonenumber" | "email"}
  * @description 
  **/
setType (value) { this[`type`] = value; return this; } 
	 
		/**
  * @type {string}
  * @description 
  **/
 password;
		/**
  * @returns {string}
  * @description 
  **/
getPassword () { return this[`password`] }
		/**
  * @param {string}
  * @description 
  **/
setPassword (value) { this[`password`] = value; return this; } 
	 
		/**
  * @type {string}
  * @description 
  **/
 firstName;
		/**
  * @returns {string}
  * @description 
  **/
getFirstName () { return this[`firstName`] }
		/**
  * @param {string}
  * @description 
  **/
setFirstName (value) { this[`firstName`] = value; return this; } 
	 
		/**
  * @type {string}
  * @description 
  **/
 lastName;
		/**
  * @returns {string}
  * @description 
  **/
getLastName () { return this[`lastName`] }
		/**
  * @param {string}
  * @description 
  **/
setLastName (value) { this[`lastName`] = value; return this; } 
	 
		/**
  * @type {string}
  * @description 
  **/
 inviteId;
		/**
  * @returns {string}
  * @description 
  **/
getInviteId () { return this[`inviteId`] }
		/**
  * @param {string}
  * @description 
  **/
setInviteId (value) { this[`inviteId`] = value; return this; } 
	 
		/**
  * @type {string}
  * @description 
  **/
 publicJoinKeyId;
		/**
  * @returns {string}
  * @description 
  **/
getPublicJoinKeyId () { return this[`publicJoinKeyId`] }
		/**
  * @param {string}
  * @description 
  **/
setPublicJoinKeyId (value) { this[`publicJoinKeyId`] = value; return this; } 
	 
		/**
  * @type {string}
  * @description 
  **/
 workspaceTypeId;
		/**
  * @returns {string}
  * @description 
  **/
getWorkspaceTypeId () { return this[`workspaceTypeId`] }
		/**
  * @param {string}
  * @description 
  **/
setWorkspaceTypeId (value) { this[`workspaceTypeId`] = value; return this; } 
	

	

	/** a placeholder for WebRequestX auto patching the json content to the object **/
	static __jsonParsable;

	
		/**
   * Nest.js decorator for controller headers. Instead of using @Headers() value: any, now you can use for example:
   * @example
   * @Get()
   * getHello(@ClassicSignupReq.Nest() headers: ClassicSignupReq): string {
   *  return JSON.stringify(headers.getContentType());
   * }
   */
  static Nest = createParamDecorator(
	(_data, ctx) => {
		// @ts-ignore
		const request = ctx.switchToHttp().getRequest();
		// @ts-ignore
		return new ClassicSignupReq( request.body );
	},
  );


	
}













	
/**
  * @decription The base class definition for classicSignupRes
  **/

export class ClassicSignupRes {
	 
		/**
  * @type {UserSessionDto}
  * @description Returns the user session in case that signup is completely successful.
  **/
 session;
		/**
  * @returns {UserSessionDto}
  * @description Returns the user session in case that signup is completely successful.
  **/
getSession () { return this[`session`] }
		/**
  * @param {UserSessionDto}
  * @description Returns the user session in case that signup is completely successful.
  **/
setSession (value) { this[`session`] = value; return this; } 
	 
		/**
  * @type {string}
  * @description If time based otp is available, we add it response to make it easier for ui.
  **/
 totpUrl;
		/**
  * @returns {string}
  * @description If time based otp is available, we add it response to make it easier for ui.
  **/
getTotpUrl () { return this[`totpUrl`] }
		/**
  * @param {string}
  * @description If time based otp is available, we add it response to make it easier for ui.
  **/
setTotpUrl (value) { this[`totpUrl`] = value; return this; } 
	 
		/**
  * @type {boolean}
  * @description Returns true and session will be empty if, the totp is required by the installation. In such scenario, you need to forward user to setup totp screen.
  **/
 continueToTotp;
		/**
  * @returns {boolean}
  * @description Returns true and session will be empty if, the totp is required by the installation. In such scenario, you need to forward user to setup totp screen.
  **/
getContinueToTotp () { return this[`continueToTotp`] }
		/**
  * @param {boolean}
  * @description Returns true and session will be empty if, the totp is required by the installation. In such scenario, you need to forward user to setup totp screen.
  **/
setContinueToTotp (value) { this[`continueToTotp`] = value; return this; } 
	 
		/**
  * @type {boolean}
  * @description Determines if user must complete totp in order to continue based on workspace or installation
  **/
 forcedTotp;
		/**
  * @returns {boolean}
  * @description Determines if user must complete totp in order to continue based on workspace or installation
  **/
getForcedTotp () { return this[`forcedTotp`] }
		/**
  * @param {boolean}
  * @description Determines if user must complete totp in order to continue based on workspace or installation
  **/
setForcedTotp (value) { this[`forcedTotp`] = value; return this; } 
	

	

	/** a placeholder for WebRequestX auto patching the json content to the object **/
	static __jsonParsable;

	
		/**
   * Nest.js decorator for controller headers. Instead of using @Headers() value: any, now you can use for example:
   * @example
   * @Get()
   * getHello(@ClassicSignupRes.Nest() headers: ClassicSignupRes): string {
   *  return JSON.stringify(headers.getContentType());
   * }
   */
  static Nest = createParamDecorator(
	(_data, ctx) => {
		// @ts-ignore
		const request = ctx.switchToHttp().getRequest();
		// @ts-ignore
		return new ClassicSignupRes( request.body );
	},
  );


	
}








/**
 * ClassicSignupHeaders class
 * Auto-generated from Module3Action
 */
export class ClassicSignupHeaders extends Headers {

  // the getters generated by us would be casting types before returning.
  // you still can use .get function to get the string value.
  #getTyped(key, type) {
    const val = this.get(key);
    if (val == null) return null;

    const t = type.toLowerCase();

    if (t.includes('number')) return Number(val);
    if (t.includes('bool')) return val === 'true';
    return val; // string or any other fallback
  }


  /**
   * @returns {Record<string, string>}
   * Converts Headers to plain object
   */
  toObject() {
    return Object.fromEntries(this.entries());
  }

  
  /**
   * Nest.js decorator for controller headers. Instead of using @Headers() value: any, now you can use for example:
   * @example
   * @Get()
   * getHello(@ClassicSignupHeaders.Nest() headers: ClassicSignupHeaders): string {
   *  return JSON.stringify(headers.getContentType());
   * }
   */
  static Nest = createParamDecorator(
	(_data, ctx) => {
		// @ts-ignore
		const request = ctx.switchToHttp().getRequest();
		// @ts-ignore
		return new ClassicSignupHeaders(Object.entries(request.headers));
	},
  );

  
}




/**
 * ClassicSignupQueryParams class
 * Auto-generated from Module3Action
 */
export class ClassicSignupQueryParams extends URLSearchParamsX {


  // the getters generated by us would be casting types before returning.
  // you still can use .get function to get the string value.
  #getTyped(key, type) {
    const val = this.get(key);
    if (val == null) return null;

    const t = type.toLowerCase();

    if (t.includes('number')) return Number(val);
    if (t.includes('bool')) return val === 'true';
    return val; // string or any other fallback
  }


  
  /**
   * Nest.js decorator for controller query. Instead of using @Query() value: any, now you can use for example:
   * @example
   * @Get()
   * getHello(@ClassicSignupQueryParams.Nest() query: ClassicSignupQueryParams): string {
   *  return JSON.stringify(query.getMyfield());
   * }
   */
  static Nest = createParamDecorator(
	(_data, ctx) => {
		// @ts-ignore
		const request = ctx.switchToHttp().getRequest();
		// @ts-ignore
		return new ClassicSignupQueryParams(request.query);
	},
  );

  
}


 
