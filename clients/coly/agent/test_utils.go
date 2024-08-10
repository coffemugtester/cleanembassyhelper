package agent

// TODO: actual html files to load
// TODO: what's the actual name of test utils
const noEmbassiesFound = `
<html>
	<head>
	</head>
	<body>
		<div>Hello World!</div>
	</body>
</html>
`

const spainEmbassiesInGermanyHTML = `
<html>
<div class="embassy__wrapper embassy__wrapper--hide-ads">
                    <div class="embassy_view_main">
                        <div class="embassy_view_left col-sm-12 alignc">
                            <div class="clearr"></div>
                            <div class="embassy-registration-widget-vertical" id="embreg-widget"></div>
                        </div>
                        <div class="embassy_view_center col-sm-12 customs_h3 float_left">
                            <div class="embassy__list">
                                <h2 class="embassy__name">Spain Embassy 
 in Berlin</h2>
                                <div class="embassy__list__wrap">
                                    <div class="embassy__list__body">
                                        <div class="flex-row embassy__list__body-flex">
                                            <div class="embassy__list__address vcard">
                                                <div class="embassy__description">
                                                    <div class="embassy__address__info embassy__address clearfix">
                                                        <div class="vi__icon">
                                                            <i class="embassy__address-icon iconfont iconfont--marker"></i>
                                                        </div>
                                                        <div class="vi__icon_text">
                                                            <b>Address</b>
                                                            <div>
                                                                <span class="adr">
                                                                    Lichtensteinallee, 1<br>
                                                                    10787<br>
                                                                    Berlin<br>Germany
                                                                </span>
                                                            </div>
                                                        </div>
                                                    </div>
                                                    <div class="embassy__address__info clearfix embassy__phone ">
                                                        <div class="vi__icon">
                                                            <i class="embassy__phone-icon iconfont iconfont--phone"></i>
                                                        </div>
                                                        <div class="vi__icon_text">
                                                            <b>Phone</b>
                                                            <div>
                                                                <strong class="tel">
                                                                    <span class="d-inline-block" x-ms-format-detection="none">+49-30-254-0070</span>
                                                                </strong>
                                                            </div>
                                                        </div>
                                                    </div>
                                                    <div class="embassy__address__info clearfix ">
                                                        <div class="vi__icon">
                                                            <i class="embassy__fax-icon iconfont iconfont--fax"></i>
                                                        </div>
                                                        <div class="vi__icon_text">
                                                            <b>Fax</b>
                                                            <div>
                                                                <strong class="fax">
                                                                    <span class="d-inline-block" x-ms-format-detection="none">+49-30-257-99557</span>
                                                                </strong>
                                                            </div>
                                                        </div>
                                                    </div>
                                                    <div class="embassy__address__info clearfix embassy__email ">
                                                        <div class="vi__icon"></div>
                                                        <div class="vi__icon_text">
                                                            <b>Email</b>
                                                            <div>
                                                                <strong class="email">Emb.Berlin.inf@maec.es</strong>
                                                            </div>
                                                        </div>
                                                    </div>
                                                    <div class="embassy__address__info clearfix embassy__site ">
                                                        <div class="vi__icon"></div>
                                                        <div class="vi__icon_text">
                                                            <b>Website URL</b>
                                                            <div>
                                                                <strong class="url">http://www.exteriores.gob.es/Embajadas/Berlin/es/Paginas/inicio.aspx</strong>
                                                            </div>
                                                        </div>
                                                    </div>
                                                </div>
                                            </div>
                                            <div class="embassy__map">
                                                <a href="https://www.google.com/maps?q=Spain+Embassy+
+in+Berlin+Lichtensteinallee%2C+1+10787+Berlin+Germany" class="embassy__map-link" rel="nofollow noopener" target="_blank">
                                                    <span class="embassy__map-link-text">Map link</span>
                                                </a>
                                            </div>
                                        </div>
                                    </div>
                                    <a href="" class="notify_us" id="notify_us_18634">Report changes</a>
                                    <div class="embassy__report" id="sbt_report_18634">
                                        <div class="embassy__report-close" id="sbt_close_18634">&#215;</div>
                                        <p class="embassy__report-changes">Report changes</p>
                                        <div class="alert alert-error" style="display:none">
                                            <button class="close btn-close-npost-status" type="button">&#215;</button>
                                            <div class="embassy__report-result"></div>
                                        </div>
                                        <form method="post" class="form_container_new">
                                            <textarea name="npost_content" class="npost_content" placeholder="Your message"></textarea>
                                            <div class="flex-row embassy__report-flex">
                                                <div class="embassy__report-recaptcha">
                                                    <div id="RecaptchaField_18634" class="RecaptchaField"></div>
                                                </div>
                                                <div class="embassy__report-submit">
                                                    <input type="hidden" name="action" value="add_post">
                                                    <input type="hidden" name="currentaddress" value="Lichtensteinallee, 1&lt;br&gt;10787&lt;br&gt;Berlin&lt;br&gt;Germany">
                                                    <input type="submit" value="Submit" class="btn btn-success embassy__report-btn">
                                                </div>
                                            </div>
                                        </form>
                                    </div>
                                </div>
                                <h2 class="embassy__name">Spain Consulate 
 in Düsseldorf</h2>
                                <div class="embassy__list__wrap">
                                    <div class="embassy__list__body">
                                        <div class="flex-row embassy__list__body-flex">
                                            <div class="embassy__list__address vcard">
                                                <div class="embassy__description">
                                                    <div class="embassy__address__info embassy__address clearfix">
                                                        <div class="vi__icon">
                                                            <i class="embassy__address-icon iconfont iconfont--marker"></i>
                                                        </div>
                                                        <div class="vi__icon_text">
                                                            <b>Address</b>
                                                            <div>
                                                                <span class="adr">
                                                                    Hombergerstr., 16<br>
                                                                    40474<br>
                                                                    Düsseldorf<br>Germany
                                                                </span>
                                                            </div>
                                                        </div>
                                                    </div>
                                                    <div class="embassy__address__info clearfix embassy__phone ">
                                                        <div class="vi__icon">
                                                            <i class="embassy__phone-icon iconfont iconfont--phone"></i>
                                                        </div>
                                                        <div class="vi__icon_text">
                                                            <b>Phone</b>
                                                            <div>
                                                                <strong class="tel">
                                                                    <span class="d-inline-block" x-ms-format-detection="none">+49-211-439-080</span>
                                                                </strong>
                                                            </div>
                                                        </div>
                                                    </div>
                                                    <div class="embassy__address__info clearfix ">
                                                        <div class="vi__icon">
                                                            <i class="embassy__fax-icon iconfont iconfont--fax"></i>
                                                        </div>
                                                        <div class="vi__icon_text">
                                                            <b>Fax</b>
                                                            <div>
                                                                <strong class="fax">
                                                                    <span class="d-inline-block" x-ms-format-detection="none">+49-211-438-0315</span>
                                                                </strong>
                                                            </div>
                                                        </div>
                                                    </div>
                                                    <div class="embassy__address__info clearfix embassy__email ">
                                                        <div class="vi__icon"></div>
                                                        <div class="vi__icon_text">
                                                            <b>Email</b>
                                                            <div>
                                                                <strong class="email">
                                                                    cog.dusseldorf@maec.es<br>cog.dusseldorf.vis@maec.es
                                                                </strong>
                                                            </div>
                                                        </div>
                                                    </div>
                                                </div>
                                            </div>
                                            <div class="embassy__map">
                                                <a href="https://www.google.com/maps?q=Spain+Consulate+
+in+Düsseldorf+Hombergerstr.%2C+16+40474+D%C3%BCsseldorf+Germany" class="embassy__map-link" rel="nofollow noopener" target="_blank">
                                                    <span class="embassy__map-link-text">Map link</span>
                                                </a>
                                            </div>
                                        </div>
                                    </div>
                                    <a href="" class="notify_us" id="notify_us_18636">Report changes</a>
                                    <div class="embassy__report" id="sbt_report_18636">
                                        <div class="embassy__report-close" id="sbt_close_18636">&#215;</div>
                                        <p class="embassy__report-changes">Report changes</p>
                                        <div class="alert alert-error" style="display:none">
                                            <button class="close btn-close-npost-status" type="button">&#215;</button>
                                            <div class="embassy__report-result"></div>
                                        </div>
                                        <form method="post" class="form_container_new">
                                            <textarea name="npost_content" class="npost_content" placeholder="Your message"></textarea>
                                            <div class="flex-row embassy__report-flex">
                                                <div class="embassy__report-recaptcha">
                                                    <div id="RecaptchaField_18636" class="RecaptchaField"></div>
                                                </div>
                                                <div class="embassy__report-submit">
                                                    <input type="hidden" name="action" value="add_post">
                                                    <input type="hidden" name="currentaddress" value="Hombergerstr., 16&lt;br&gt;40474&lt;br&gt;Düsseldorf&lt;br&gt;Germany">
                                                    <input type="submit" value="Submit" class="btn btn-success embassy__report-btn">
                                                </div>
                                            </div>
                                        </form>
                                    </div>
                                </div>
                                <h2 class="embassy__name">Spain Consulate 
 in Stuttgart</h2>
                                <div class="embassy__list__wrap">
                                    <div class="embassy__list__body">
                                        <div class="flex-row embassy__list__body-flex">
                                            <div class="embassy__list__address vcard">
                                                <div class="embassy__description">
                                                    <div class="embassy__address__info embassy__address clearfix">
                                                        <div class="vi__icon">
                                                            <i class="embassy__address-icon iconfont iconfont--marker"></i>
                                                        </div>
                                                        <div class="vi__icon_text">
                                                            <b>Address</b>
                                                            <div>
                                                                <span class="adr">
                                                                    Lenzhalde, 61<br>
                                                                    70192<br>
                                                                    Stuttgart<br>Germany
                                                                </span>
                                                            </div>
                                                        </div>
                                                    </div>
                                                    <div class="embassy__address__info clearfix embassy__phone ">
                                                        <div class="vi__icon">
                                                            <i class="embassy__phone-icon iconfont iconfont--phone"></i>
                                                        </div>
                                                        <div class="vi__icon_text">
                                                            <b>Phone</b>
                                                            <div>
                                                                <strong class="tel">
                                                                    <span class="d-inline-block" x-ms-format-detection="none">+49-711-997-9800</span>
                                                                </strong>
                                                            </div>
                                                        </div>
                                                    </div>
                                                    <div class="embassy__address__info clearfix ">
                                                        <div class="vi__icon">
                                                            <i class="embassy__fax-icon iconfont iconfont--fax"></i>
                                                        </div>
                                                        <div class="vi__icon_text">
                                                            <b>Fax</b>
                                                            <div>
                                                                <strong class="fax">
                                                                    <span class="d-inline-block" x-ms-format-detection="none">+49-711-226-5927</span>
                                                                </strong>
                                                            </div>
                                                        </div>
                                                    </div>
                                                </div>
                                            </div>
                                            <div class="embassy__map">
                                                <a href="https://www.google.com/maps?q=Spain+Consulate+
+in+Stuttgart+Lenzhalde%2C+61+70192+Stuttgart+Germany" class="embassy__map-link" rel="nofollow noopener" target="_blank">
                                                    <span class="embassy__map-link-text">Map link</span>
                                                </a>
                                            </div>
                                        </div>
                                    </div>
                                    <a href="" class="notify_us" id="notify_us_18637">Report changes</a>
                                    <div class="embassy__report" id="sbt_report_18637">
                                        <div class="embassy__report-close" id="sbt_close_18637">&#215;</div>
                                        <p class="embassy__report-changes">Report changes</p>
                                        <div class="alert alert-error" style="display:none">
                                            <button class="close btn-close-npost-status" type="button">&#215;</button>
                                            <div class="embassy__report-result"></div>
                                        </div>
                                        <form method="post" class="form_container_new">
                                            <textarea name="npost_content" class="npost_content" placeholder="Your message"></textarea>
                                            <div class="flex-row embassy__report-flex">
                                                <div class="embassy__report-recaptcha">
                                                    <div id="RecaptchaField_18637" class="RecaptchaField"></div>
                                                </div>
                                                <div class="embassy__report-submit">
                                                    <input type="hidden" name="action" value="add_post">
                                                    <input type="hidden" name="currentaddress" value="Lenzhalde, 61&lt;br&gt;70192&lt;br&gt;Stuttgart&lt;br&gt;Germany">
                                                    <input type="submit" value="Submit" class="btn btn-success embassy__report-btn">
                                                </div>
                                            </div>
                                        </form>
                                    </div>
                                </div>
                                <h2 class="embassy__name">Spain Consulate 
 in Hamburg</h2>
                                <div class="embassy__list__wrap">
                                    <div class="embassy__list__body">
                                        <div class="flex-row embassy__list__body-flex">
                                            <div class="embassy__list__address vcard">
                                                <div class="embassy__description">
                                                    <div class="embassy__address__info embassy__address clearfix">
                                                        <div class="vi__icon">
                                                            <i class="embassy__address-icon iconfont iconfont--marker"></i>
                                                        </div>
                                                        <div class="vi__icon_text">
                                                            <b>Address</b>
                                                            <div>
                                                                <span class="adr">
                                                                    Mittelweg, 37<br>
                                                                    20148<br>
                                                                    Hamburg<br>Germany
                                                                </span>
                                                            </div>
                                                        </div>
                                                    </div>
                                                    <div class="embassy__address__info clearfix embassy__phone ">
                                                        <div class="vi__icon">
                                                            <i class="embassy__phone-icon iconfont iconfont--phone"></i>
                                                        </div>
                                                        <div class="vi__icon_text">
                                                            <b>Phone</b>
                                                            <div>
                                                                <strong class="tel">
                                                                    <span class="d-inline-block" x-ms-format-detection="none">
                                                                        +49-40-414-646<br>+49-40-414-640
                                                                    </span>
                                                                </strong>
                                                            </div>
                                                        </div>
                                                    </div>
                                                    <div class="embassy__address__info clearfix ">
                                                        <div class="vi__icon">
                                                            <i class="embassy__fax-icon iconfont iconfont--fax"></i>
                                                        </div>
                                                        <div class="vi__icon_text">
                                                            <b>Fax</b>
                                                            <div>
                                                                <strong class="fax">
                                                                    <span class="d-inline-block" x-ms-format-detection="none">+49-40-417-449</span>
                                                                </strong>
                                                            </div>
                                                        </div>
                                                    </div>
                                                    <div class="embassy__address__info clearfix embassy__email ">
                                                        <div class="vi__icon"></div>
                                                        <div class="vi__icon_text">
                                                            <b>Email</b>
                                                            <div>
                                                                <strong class="email">cog.hamburgo@maec.es</strong>
                                                            </div>
                                                        </div>
                                                    </div>
                                                    <div class="embassy__address__info clearfix embassy__site ">
                                                        <div class="vi__icon"></div>
                                                        <div class="vi__icon_text">
                                                            <b>Website URL</b>
                                                            <div>
                                                                <strong class="url">www.consuladoenhamburgo.maec.es</strong>
                                                            </div>
                                                        </div>
                                                    </div>
                                                </div>
                                            </div>
                                            <div class="embassy__map">
                                                <a href="https://www.google.com/maps?q=Spain+Consulate+
+in+Hamburg+Mittelweg%2C+37+20148+Hamburg+Germany" class="embassy__map-link" rel="nofollow noopener" target="_blank">
                                                    <span class="embassy__map-link-text">Map link</span>
                                                </a>
                                            </div>
                                        </div>
                                    </div>
                                    <a href="" class="notify_us" id="notify_us_18638">Report changes</a>
                                    <div class="embassy__report" id="sbt_report_18638">
                                        <div class="embassy__report-close" id="sbt_close_18638">&#215;</div>
                                        <p class="embassy__report-changes">Report changes</p>
                                        <div class="alert alert-error" style="display:none">
                                            <button class="close btn-close-npost-status" type="button">&#215;</button>
                                            <div class="embassy__report-result"></div>
                                        </div>
                                        <form method="post" class="form_container_new">
                                            <textarea name="npost_content" class="npost_content" placeholder="Your message"></textarea>
                                            <div class="flex-row embassy__report-flex">
                                                <div class="embassy__report-recaptcha">
                                                    <div id="RecaptchaField_18638" class="RecaptchaField"></div>
                                                </div>
                                                <div class="embassy__report-submit">
                                                    <input type="hidden" name="action" value="add_post">
                                                    <input type="hidden" name="currentaddress" value="Mittelweg, 37&lt;br&gt;20148&lt;br&gt;Hamburg&lt;br&gt;Germany">
                                                    <input type="submit" value="Submit" class="btn btn-success embassy__report-btn">
                                                </div>
                                            </div>
                                        </form>
                                    </div>
                                </div>
                                <h2 class="embassy__name">Spain Consulate 
 in Frankfurt</h2>
                                <div class="embassy__list__wrap">
                                    <div class="embassy__list__body">
                                        <div class="flex-row embassy__list__body-flex">
                                            <div class="embassy__list__address vcard">
                                                <div class="embassy__description">
                                                    <div class="embassy__address__info embassy__address clearfix">
                                                        <div class="vi__icon">
                                                            <i class="embassy__address-icon iconfont iconfont--marker"></i>
                                                        </div>
                                                        <div class="vi__icon_text">
                                                            <b>Address</b>
                                                            <div>
                                                                <span class="adr">
                                                                    Nibelungenplatz, 3<br>
                                                                    60318<br>
                                                                    Frankfurt<br>Germany
                                                                </span>
                                                            </div>
                                                        </div>
                                                    </div>
                                                    <div class="embassy__address__info clearfix embassy__phone ">
                                                        <div class="vi__icon">
                                                            <i class="embassy__phone-icon iconfont iconfont--phone"></i>
                                                        </div>
                                                        <div class="vi__icon_text">
                                                            <b>Phone</b>
                                                            <div>
                                                                <strong class="tel">
                                                                    <span class="d-inline-block" x-ms-format-detection="none">+49-69-959-1660</span>
                                                                </strong>
                                                            </div>
                                                        </div>
                                                    </div>
                                                    <div class="embassy__address__info clearfix ">
                                                        <div class="vi__icon">
                                                            <i class="embassy__fax-icon iconfont iconfont--fax"></i>
                                                        </div>
                                                        <div class="vi__icon_text">
                                                            <b>Fax</b>
                                                            <div>
                                                                <strong class="fax">
                                                                    <span class="d-inline-block" x-ms-format-detection="none">+49-69-596-4742</span>
                                                                </strong>
                                                            </div>
                                                        </div>
                                                    </div>
                                                    <div class="embassy__address__info clearfix embassy__email ">
                                                        <div class="vi__icon"></div>
                                                        <div class="vi__icon_text">
                                                            <b>Email</b>
                                                            <div>
                                                                <strong class="email">cog.francfort@maec.es</strong>
                                                            </div>
                                                        </div>
                                                    </div>
                                                </div>
                                            </div>
                                            <div class="embassy__map">
                                                <a href="https://www.google.com/maps?q=Spain+Consulate+
+in+Frankfurt+Nibelungenplatz%2C+3+60318+Frankfurt+Germany" class="embassy__map-link" rel="nofollow noopener" target="_blank">
                                                    <span class="embassy__map-link-text">Map link</span>
                                                </a>
                                            </div>
                                        </div>
                                    </div>
                                    <a href="" class="notify_us" id="notify_us_18639">Report changes</a>
                                    <div class="embassy__report" id="sbt_report_18639">
                                        <div class="embassy__report-close" id="sbt_close_18639">&#215;</div>
                                        <p class="embassy__report-changes">Report changes</p>
                                        <div class="alert alert-error" style="display:none">
                                            <button class="close btn-close-npost-status" type="button">&#215;</button>
                                            <div class="embassy__report-result"></div>
                                        </div>
                                        <form method="post" class="form_container_new">
                                            <textarea name="npost_content" class="npost_content" placeholder="Your message"></textarea>
                                            <div class="flex-row embassy__report-flex">
                                                <div class="embassy__report-recaptcha">
                                                    <div id="RecaptchaField_18639" class="RecaptchaField"></div>
                                                </div>
                                                <div class="embassy__report-submit">
                                                    <input type="hidden" name="action" value="add_post">
                                                    <input type="hidden" name="currentaddress" value="Nibelungenplatz, 3&lt;br&gt;60318&lt;br&gt;Frankfurt&lt;br&gt;Germany">
                                                    <input type="submit" value="Submit" class="btn btn-success embassy__report-btn">
                                                </div>
                                            </div>
                                        </form>
                                    </div>
                                </div>
                                <h2 class="embassy__name">Spain Consulate 
 in München</h2>
                                <div class="embassy__list__wrap">
                                    <div class="embassy__list__body">
                                        <div class="flex-row embassy__list__body-flex">
                                            <div class="embassy__list__address vcard">
                                                <div class="embassy__description">
                                                    <div class="embassy__address__info embassy__address clearfix">
                                                        <div class="vi__icon">
                                                            <i class="embassy__address-icon iconfont iconfont--marker"></i>
                                                        </div>
                                                        <div class="vi__icon_text">
                                                            <b>Address</b>
                                                            <div>
                                                                <span class="adr">
                                                                    Oberföhringer Str., 45<br>
                                                                    81925<br>
                                                                    München<br>Germany
                                                                </span>
                                                            </div>
                                                        </div>
                                                    </div>
                                                    <div class="embassy__address__info clearfix embassy__phone ">
                                                        <div class="vi__icon">
                                                            <i class="embassy__phone-icon iconfont iconfont--phone"></i>
                                                        </div>
                                                        <div class="vi__icon_text">
                                                            <b>Phone</b>
                                                            <div>
                                                                <strong class="tel">
                                                                    <span class="d-inline-block" x-ms-format-detection="none">+49-89-998-4790</span>
                                                                </strong>
                                                            </div>
                                                        </div>
                                                    </div>
                                                    <div class="embassy__address__info clearfix ">
                                                        <div class="vi__icon">
                                                            <i class="embassy__fax-icon iconfont iconfont--fax"></i>
                                                        </div>
                                                        <div class="vi__icon_text">
                                                            <b>Fax</b>
                                                            <div>
                                                                <strong class="fax">
                                                                    <span class="d-inline-block" x-ms-format-detection="none">+49-89-981-0206</span>
                                                                </strong>
                                                            </div>
                                                        </div>
                                                    </div>
                                                    <div class="embassy__address__info clearfix embassy__email ">
                                                        <div class="vi__icon"></div>
                                                        <div class="vi__icon_text">
                                                            <b>Email</b>
                                                            <div>
                                                                <strong class="email">cog.munich@maec.es</strong>
                                                            </div>
                                                        </div>
                                                    </div>
                                                </div>
                                            </div>
                                            <div class="embassy__map">
                                                <a href="https://www.google.com/maps?q=Spain+Consulate+
+in+München+Oberf%C3%B6hringer+Str.%2C+45+81925+M%C3%BCnchen+Germany" class="embassy__map-link" rel="nofollow noopener" target="_blank">
                                                    <span class="embassy__map-link-text">Map link</span>
                                                </a>
                                            </div>
                                        </div>
                                    </div>
                                    <a href="" class="notify_us" id="notify_us_18640">Report changes</a>
                                    <div class="embassy__report" id="sbt_report_18640">
                                        <div class="embassy__report-close" id="sbt_close_18640">&#215;</div>
                                        <p class="embassy__report-changes">Report changes</p>
                                        <div class="alert alert-error" style="display:none">
                                            <button class="close btn-close-npost-status" type="button">&#215;</button>
                                            <div class="embassy__report-result"></div>
                                        </div>
                                        <form method="post" class="form_container_new">
                                            <textarea name="npost_content" class="npost_content" placeholder="Your message"></textarea>
                                            <div class="flex-row embassy__report-flex">
                                                <div class="embassy__report-recaptcha">
                                                    <div id="RecaptchaField_18640" class="RecaptchaField"></div>
                                                </div>
                                                <div class="embassy__report-submit">
                                                    <input type="hidden" name="action" value="add_post">
                                                    <input type="hidden" name="currentaddress" value="Oberföhringer Str., 45&lt;br&gt;81925&lt;br&gt;München&lt;br&gt;Germany">
                                                    <input type="submit" value="Submit" class="btn btn-success embassy__report-btn">
                                                </div>
                                            </div>
                                        </form>
                                    </div>
                                </div>
                                <!-- Apply for visa Block -->
                            </div>
                        </div>
                        <div class="clearr"></div>
                    </div>
                </div>
</html>
`
