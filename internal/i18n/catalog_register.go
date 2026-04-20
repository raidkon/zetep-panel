package i18n

func init() {
	en := english()
	catalog = map[Lang]map[string]string{
		EN: en,
		RU: russian(),
		ZH: mergeMaps(en, zhStrings()),
		HI: mergeMaps(en, hiStrings()),
		ES: mergeMaps(en, esStrings()),
		FR: mergeMaps(en, frStrings()),
		AR: mergeMaps(en, arStrings()),
		BN: mergeMaps(en, bnStrings()),
		PT: mergeMaps(en, ptStrings()),
		UR: mergeMaps(en, urStrings()),
	}
}
