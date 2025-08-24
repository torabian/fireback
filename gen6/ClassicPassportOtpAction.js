import { ExecutionContext, createParamDecorator } from '@nestjs/common';
import { fetchx } from 'INTERNAL_SDK_LOCATION';
/**
* Action to communicate with the action classicPassportOtp
*/


	/**
 * FetchClassicPassportOtpAction
 */

export class FetchClassicPassportOtpAction {
  static URL = '/workspace/passport/otp';
  static Method = 'post';

  
  	static Axios = (clientInstance: AxiosInstance, config: AxiosRequestConfig<unknown>) =>
		clientInstance
		.requestrequest<unknown, AxiosResponse<unknown>, unknown>(config)

		
		.then((res) => {
			return {
			...res,

			
			// if there is a output class, create instance out of it.
			data: new ClassicPassportOtpRes(res.data),
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
		

	
		
			.then((data) => new ClassicPassportOtpRes (data));
		

	
}









	
/**
  * @decription The base class definition for classicPassportOtpReq
  **/

export class ClassicPassportOtpReq {
	 
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
 otp;
		/**
  * @returns {string}
  * @description 
  **/
getOtp () { return this[`otp`] }
		/**
  * @param {string}
  * @description 
  **/
setOtp (value) { this[`otp`] = value; return this; } 
	

	

	/** a placeholder for WebRequestX auto patching the json content to the object **/
	static __jsonParsable;

	
		/**
   * Nest.js decorator for controller headers. Instead of using @Headers() value: any, now you can use for example:
   * @example
   * @Get()
   * getHello(@ClassicPassportOtpReq.Nest() headers: ClassicPassportOtpReq): string {
   *  return JSON.stringify(headers.getContentType());
   * }
   */
  static Nest = createParamDecorator(
	(_data, ctx) => {
		// @ts-ignore
		const request = ctx.switchToHttp().getRequest();
		// @ts-ignore
		return new ClassicPassportOtpReq( request.body );
	},
  );


	
}













	
/**
  * @decription The base class definition for classicPassportOtpRes
  **/

export class ClassicPassportOtpRes {
	 
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
  * @type {string}
  * @description The session secret will be used to call complete user registeration api.
  **/
 sessionSecret;
		/**
  * @returns {string}
  * @description The session secret will be used to call complete user registeration api.
  **/
getSessionSecret () { return this[`sessionSecret`] }
		/**
  * @param {string}
  * @description The session secret will be used to call complete user registeration api.
  **/
setSessionSecret (value) { this[`sessionSecret`] = value; return this; } 
	 
		/**
  * @type {boolean}
  * @description If return true, means the OTP is correct and user needs to be created before continue the authentication processs.
  **/
 continueWithCreation;
		/**
  * @returns {boolean}
  * @description If return true, means the OTP is correct and user needs to be created before continue the authentication processs.
  **/
getContinueWithCreation () { return this[`continueWithCreation`] }
		/**
  * @param {boolean}
  * @description If return true, means the OTP is correct and user needs to be created before continue the authentication processs.
  **/
setContinueWithCreation (value) { this[`continueWithCreation`] = value; return this; } 
	

	

	/** a placeholder for WebRequestX auto patching the json content to the object **/
	static __jsonParsable;

	
		/**
   * Nest.js decorator for controller headers. Instead of using @Headers() value: any, now you can use for example:
   * @example
   * @Get()
   * getHello(@ClassicPassportOtpRes.Nest() headers: ClassicPassportOtpRes): string {
   *  return JSON.stringify(headers.getContentType());
   * }
   */
  static Nest = createParamDecorator(
	(_data, ctx) => {
		// @ts-ignore
		const request = ctx.switchToHttp().getRequest();
		// @ts-ignore
		return new ClassicPassportOtpRes( request.body );
	},
  );


	
}








/**
 * ClassicPassportOtpHeaders class
 * Auto-generated from Module3Action
 */
export class ClassicPassportOtpHeaders extends Headers {

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
   * getHello(@ClassicPassportOtpHeaders.Nest() headers: ClassicPassportOtpHeaders): string {
   *  return JSON.stringify(headers.getContentType());
   * }
   */
  static Nest = createParamDecorator(
	(_data, ctx) => {
		// @ts-ignore
		const request = ctx.switchToHttp().getRequest();
		// @ts-ignore
		return new ClassicPassportOtpHeaders(Object.entries(request.headers));
	},
  );

  
}




/**
 * ClassicPassportOtpQueryParams class
 * Auto-generated from Module3Action
 */
export class ClassicPassportOtpQueryParams extends URLSearchParamsX {


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
   * getHello(@ClassicPassportOtpQueryParams.Nest() query: ClassicPassportOtpQueryParams): string {
   *  return JSON.stringify(query.getMyfield());
   * }
   */
  static Nest = createParamDecorator(
	(_data, ctx) => {
		// @ts-ignore
		const request = ctx.switchToHttp().getRequest();
		// @ts-ignore
		return new ClassicPassportOtpQueryParams(request.query);
	},
  );

  
}


 
