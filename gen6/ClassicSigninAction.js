import { ExecutionContext, createParamDecorator } from '@nestjs/common';
import { fetchx } from 'INTERNAL_SDK_LOCATION';
/**
* Action to communicate with the action classicSignin
*/


	/**
 * FetchClassicSigninAction
 */

export class FetchClassicSigninAction {
  static URL = '/passports/signin/classic';
  static Method = 'post';

  
  	static Axios = (clientInstance: AxiosInstance, config: AxiosRequestConfig<unknown>) =>
		clientInstance
		.requestrequest<unknown, AxiosResponse<unknown>, unknown>(config)

		
		.then((res) => {
			return {
			...res,

			
			// if there is a output class, create instance out of it.
			data: new ClassicSigninRes(res.data),
			};
		});
		
	
  
  
	static Fetch = (
		init?: TypedRequestInit<unknown, unknown> | undefined,
		qs?: SignoutQueryParams,
		overrideUrl?: string
	) =>
		fetchx<unknown, unknown, unknown>(
			new URL((overrideUrl ?? FetchSignoutAction.URL) + '?' + qs?.toString()),
			init
		)

		
			.then((res) => res.json())
		

	
		
			.then((data) => new ClassicSigninRes (data));
		

	
}









	
/**
  * @decription The base class definition for classicSigninReq
  **/

export class ClassicSigninReq {
	 
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
  * @description Accepts login with totp code. If enabled, first login would return a success response with next[enter-totp] value and ui can understand that user needs to be navigated into the screen other screen.
  **/
 totpCode;
		/**
  * @returns {string}
  * @description Accepts login with totp code. If enabled, first login would return a success response with next[enter-totp] value and ui can understand that user needs to be navigated into the screen other screen.
  **/
getTotpCode () { return this[`totpCode`] }
		/**
  * @param {string}
  * @description Accepts login with totp code. If enabled, first login would return a success response with next[enter-totp] value and ui can understand that user needs to be navigated into the screen other screen.
  **/
setTotpCode (value) { this[`totpCode`] = value; return this; } 
	 
		/**
  * @type {string}
  * @description Session secret when logging in to the application requires more steps to complete.
  **/
 sessionSecret;
		/**
  * @returns {string}
  * @description Session secret when logging in to the application requires more steps to complete.
  **/
getSessionSecret () { return this[`sessionSecret`] }
		/**
  * @param {string}
  * @description Session secret when logging in to the application requires more steps to complete.
  **/
setSessionSecret (value) { this[`sessionSecret`] = value; return this; } 
	

	

	/** a placeholder for WebRequestX auto patching the json content to the object **/
	static __jsonParsable;

	
		/**
   * Nest.js decorator for controller headers. Instead of using @Headers() value: any, now you can use for example:
   * @example
   * @Get()
   * getHello(@ClassicSigninReq.Nest() headers: ClassicSigninReq): string {
   *  return JSON.stringify(headers.getContentType());
   * }
   */
  static Nest = createParamDecorator(
	(_data, ctx) => {
		// @ts-ignore
		const request = ctx.switchToHttp().getRequest();
		// @ts-ignore
		return new ClassicSigninReq( request.body );
	},
  );


	
}













	
/**
  * @decription The base class definition for classicSigninRes
  **/

export class ClassicSigninRes {
	 
		/**
  * @type {UserSessionDto}
  * @description 
  **/
 session;
		/**
  * @returns {UserSessionDto}
  * @description 
  **/
getSession () { return this[`session`] }
		/**
  * @param {UserSessionDto}
  * @description 
  **/
setSession (value) { this[`session`] = value; return this; } 
	 
		/**
  * @type {string[]}
  * @description The next possible action which is suggested.
  **/
 next;
		/**
  * @returns {string[]}
  * @description The next possible action which is suggested.
  **/
getNext () { return this[`next`] }
		/**
  * @param {string[]}
  * @description The next possible action which is suggested.
  **/
setNext (value) { this[`next`] = value; return this; } 
	 
		/**
  * @type {string}
  * @description In case the account doesn't have totp, but enforced by installation, this value will contain the link
  **/
 totpUrl;
		/**
  * @returns {string}
  * @description In case the account doesn't have totp, but enforced by installation, this value will contain the link
  **/
getTotpUrl () { return this[`totpUrl`] }
		/**
  * @param {string}
  * @description In case the account doesn't have totp, but enforced by installation, this value will contain the link
  **/
setTotpUrl (value) { this[`totpUrl`] = value; return this; } 
	 
		/**
  * @type {string}
  * @description Returns a secret session if the authentication requires more steps.
  **/
 sessionSecret;
		/**
  * @returns {string}
  * @description Returns a secret session if the authentication requires more steps.
  **/
getSessionSecret () { return this[`sessionSecret`] }
		/**
  * @param {string}
  * @description Returns a secret session if the authentication requires more steps.
  **/
setSessionSecret (value) { this[`sessionSecret`] = value; return this; } 
	

	

	/** a placeholder for WebRequestX auto patching the json content to the object **/
	static __jsonParsable;

	
		/**
   * Nest.js decorator for controller headers. Instead of using @Headers() value: any, now you can use for example:
   * @example
   * @Get()
   * getHello(@ClassicSigninRes.Nest() headers: ClassicSigninRes): string {
   *  return JSON.stringify(headers.getContentType());
   * }
   */
  static Nest = createParamDecorator(
	(_data, ctx) => {
		// @ts-ignore
		const request = ctx.switchToHttp().getRequest();
		// @ts-ignore
		return new ClassicSigninRes( request.body );
	},
  );


	
}








/**
 * ClassicSigninHeaders class
 * Auto-generated from Module3Action
 */
export class ClassicSigninHeaders extends Headers {

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
   * getHello(@ClassicSigninHeaders.Nest() headers: ClassicSigninHeaders): string {
   *  return JSON.stringify(headers.getContentType());
   * }
   */
  static Nest = createParamDecorator(
	(_data, ctx) => {
		// @ts-ignore
		const request = ctx.switchToHttp().getRequest();
		// @ts-ignore
		return new ClassicSigninHeaders(Object.entries(request.headers));
	},
  );

  
}




/**
 * ClassicSigninQueryParams class
 * Auto-generated from Module3Action
 */
export class ClassicSigninQueryParams extends URLSearchParamsX {


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
   * getHello(@ClassicSigninQueryParams.Nest() query: ClassicSigninQueryParams): string {
   *  return JSON.stringify(query.getMyfield());
   * }
   */
  static Nest = createParamDecorator(
	(_data, ctx) => {
		// @ts-ignore
		const request = ctx.switchToHttp().getRequest();
		// @ts-ignore
		return new ClassicSigninQueryParams(request.query);
	},
  );

  
}


 
