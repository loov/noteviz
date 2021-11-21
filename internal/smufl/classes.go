package smufl

const (
	// 4-string tab clef
	FourstringTabClef = '\uE06E'
	// 6-string tab clef
	SixstringTabClef = '\uE06D'
	// 11 large diesis down, 3° down [46 EDO]
	AccSagittal11LargeDiesisDown = '\uE30D'
	// 11 large diesis up, (11L), (sharp less 11M), 3° up [46 EDO]
	AccSagittal11LargeDiesisUp = '\uE30C'
	// 11 medium diesis down, 1°[17 31] 2°46 down, 1/4-tone down
	AccSagittal11MediumDiesisDown = '\uE30B'
	// 11 medium diesis up, (11M), 1°[17 31] 2°46 up, 1/4-tone up
	AccSagittal11MediumDiesisUp = '\uE30A'
	// 11:19 large diesis down
	AccSagittal11v19LargeDiesisDown = '\uE3AB'
	// 11:19 large diesis up, (11:19L, apotome less 11:19M)
	AccSagittal11v19LargeDiesisUp = '\uE3AA'
	// 11:19 medium diesis down
	AccSagittal11v19MediumDiesisDown = '\uE3A3'
	// 11:19 medium diesis up, (11:19M, 11M plus 19s)
	AccSagittal11v19MediumDiesisUp = '\uE3A2'
	// 11:49 comma down
	AccSagittal11v49CommaDown = '\uE397'
	// 11:49 comma up, (11:49C, 11M less 49C)
	AccSagittal11v49CommaUp = '\uE396'
	// 143 comma down
	AccSagittal143CommaDown = '\uE395'
	// 143 comma up, (143C, 13L less 11M)
	AccSagittal143CommaUp = '\uE394'
	// 17 comma down
	AccSagittal17CommaDown = '\uE343'
	// 17 comma up, (17C)
	AccSagittal17CommaUp = '\uE342'
	// 17 kleisma down
	AccSagittal17KleismaDown = '\uE393'
	// 17 kleisma up, (17k)
	AccSagittal17KleismaUp = '\uE392'
	// 19 comma down
	AccSagittal19CommaDown = '\uE399'
	// 19 comma up, (19C)
	AccSagittal19CommaUp = '\uE398'
	// 19 schisma down
	AccSagittal19SchismaDown = '\uE391'
	// 19 schisma up, (19s)
	AccSagittal19SchismaUp = '\uE390'
	// 1 mina down, 1/(5⋅7⋅13)-schismina down, 0.42 cents down
	AccSagittal1MinaDown = '\uE3F5'
	// 1 mina up, 1/(5⋅7⋅13)-schismina up, 0.42 cents up
	AccSagittal1MinaUp = '\uE3F4'
	// 1 tina down, 7²⋅11⋅19/5-schismina down, 0.17 cents down
	AccSagittal1TinaDown = '\uE3F9'
	// 1 tina up, 7²⋅11⋅19/5-schismina up, 0.17 cents up
	AccSagittal1TinaUp = '\uE3F8'
	// 23 comma down, 2° down [96 EDO], 1/8-tone down
	AccSagittal23CommaDown = '\uE371'
	// 23 comma up, (23C), 2° up [96 EDO], 1/8-tone up
	AccSagittal23CommaUp = '\uE370'
	// 23 small diesis down
	AccSagittal23SmallDiesisDown = '\uE39F'
	// 23 small diesis up, (23S)
	AccSagittal23SmallDiesisUp = '\uE39E'
	// 25 small diesis down, 2° down [53 EDO]
	AccSagittal25SmallDiesisDown = '\uE307'
	// 25 small diesis up, (25S, ~5:13S, ~37S, 5C plus 5C), 2° up [53 EDO]
	AccSagittal25SmallDiesisUp = '\uE306'
	// 2 minas down, 65/77-schismina down, 0.83 cents down
	AccSagittal2MinasDown = '\uE3F7'
	// 2 minas up, 65/77-schismina up, 0.83 cents up
	AccSagittal2MinasUp = '\uE3F6'
	// 2 tinas down, 1/(7³⋅17)-schismina down, 0.30 cents down
	AccSagittal2TinasDown = '\uE3FB'
	// 2 tinas up, 1/(7³⋅17)-schismina up, 0.30 cents up
	AccSagittal2TinasUp = '\uE3FA'
	// 35 large diesis down, 2° down [50 EDO], 5/18-tone down
	AccSagittal35LargeDiesisDown = '\uE30F'
	// 35 large diesis up, (35L, ~13L, ~125L, sharp less 35M), 2°50 up
	AccSagittal35LargeDiesisUp = '\uE30E'
	// 35 medium diesis down, 1°[50] 2°[27] down, 2/9-tone down
	AccSagittal35MediumDiesisDown = '\uE309'
	// 35 medium diesis up, (35M, ~13M, ~125M, 5C plus 7C), 2/9-tone up
	AccSagittal35MediumDiesisUp = '\uE308'
	// 3 tinas down, 1 mina down, 1/(5⋅7⋅13)-schismina down, 0.42 cents down
	AccSagittal3TinasDown = '\uE3FD'
	// 3 tinas up, 1 mina up, 1/(5⋅7⋅13)-schismina up, 0.42 cents up
	AccSagittal3TinasUp = '\uE3FC'
	// 49 large diesis down
	AccSagittal49LargeDiesisDown = '\uE3A9'
	// 49 large diesis up, (49L, ~31L, apotome less 49M)
	AccSagittal49LargeDiesisUp = '\uE3A8'
	// 49 medium diesis down
	AccSagittal49MediumDiesisDown = '\uE3A5'
	// 49 medium diesis up, (49M, ~31M, 7C plus 7C)
	AccSagittal49MediumDiesisUp = '\uE3A4'
	// 49 small diesis down
	AccSagittal49SmallDiesisDown = '\uE39D'
	// 49 small diesis up, (49S, ~31S)
	AccSagittal49SmallDiesisUp = '\uE39C'
	// 4 tinas down, 5²⋅11²/7-schismina down, 0.57 cents down
	AccSagittal4TinasDown = '\uE3FF'
	// 4 tinas up, 5²⋅11²/7-schismina up, 0.57 cents up
	AccSagittal4TinasUp = '\uE3FE'
	// 55 comma down, 3° down [96 EDO], 3/16-tone down
	AccSagittal55CommaDown = '\uE345'
	// 55 comma up, (55C, 11M less 5C), 3°up [96 EDO], 3/16-tone up
	AccSagittal55CommaUp = '\uE344'
	// 5 comma down, 1° down [22 27 29 34 41 46 53 96 EDOs], 1/12-tone down
	AccSagittal5CommaDown = '\uE303'
	// 5 comma up, (5C), 1° up [22 27 29 34 41 46 53 96 EDOs], 1/12-tone up
	AccSagittal5CommaUp = '\uE302'
	// 5 tinas down, 7⁴/25-schismina down, 0.72 cents down
	AccSagittal5TinasDown = '\uE401'
	// 5 tinas up, 7⁴/25-schismina up, 0.72 cents up
	AccSagittal5TinasUp = '\uE400'
	// 5:11 small diesis down
	AccSagittal5v11SmallDiesisDown = '\uE349'
	// 5:11 small diesis up, (5:11S, ~7:13S, ~11:17S, 5:7k plus 7:11C)
	AccSagittal5v11SmallDiesisUp = '\uE348'
	// 5:13 large diesis down
	AccSagittal5v13LargeDiesisDown = '\uE3AD'
	// 5:13 large diesis up, (5:13L, ~37L, apotome less 5:13M)
	AccSagittal5v13LargeDiesisUp = '\uE3AC'
	// 5:13 medium diesis down
	AccSagittal5v13MediumDiesisDown = '\uE3A1'
	// 5:13 medium diesis up, (5:13M, ~37M, 5C plus 13C)
	AccSagittal5v13MediumDiesisUp = '\uE3A0'
	// 5:19 comma down, 1/20-tone down
	AccSagittal5v19CommaDown = '\uE373'
	// 5:19 comma up, (5:19C, 5C plus 19s), 1/20-tone up
	AccSagittal5v19CommaUp = '\uE372'
	// 5:23 small diesis down, 2° down [60 EDO], 1/5-tone down
	AccSagittal5v23SmallDiesisDown = '\uE375'
	// 5:23 small diesis up, (5:23S, 5C plus 23C), 2° up [60 EDO], 1/5-tone up
	AccSagittal5v23SmallDiesisUp = '\uE374'
	// 5:49 medium diesis down
	AccSagittal5v49MediumDiesisDown = '\uE3A7'
	// 5:49 medium diesis up, (5:49M, half apotome)
	AccSagittal5v49MediumDiesisUp = '\uE3A6'
	// 5:7 kleisma down
	AccSagittal5v7KleismaDown = '\uE301'
	// 5:7 kleisma up, (5:7k, ~11:13k, 7C less 5C)
	AccSagittal5v7KleismaUp = '\uE300'
	// 6 tinas down, 2 minas down, 65/77-schismina down, 0.83 cents down
	AccSagittal6TinasDown = '\uE403'
	// 6 tinas up, 2 minas up, 65/77-schismina up, 0.83 cents up
	AccSagittal6TinasUp = '\uE402'
	// 7 comma down, 1° down [43 EDO], 2° down [72 EDO], 1/6-tone down
	AccSagittal7CommaDown = '\uE305'
	// 7 comma up, (7C), 1° up [43 EDO], 2° up [72 EDO], 1/6-tone up
	AccSagittal7CommaUp = '\uE304'
	// 7 tinas down, 7/(5²⋅17)-schismina down, 1.02 cents down
	AccSagittal7TinasDown = '\uE405'
	// 7 tinas up, 7/(5²⋅17)-schismina up, 1.02 cents up
	AccSagittal7TinasUp = '\uE404'
	// 7:11 comma down, 1° down [60 EDO], 1/10-tone down
	AccSagittal7v11CommaDown = '\uE347'
	// 7:11 comma up, (7:11C, ~13:17S, ~29S, 11L less 7C), 1° up [60 EDO]
	AccSagittal7v11CommaUp = '\uE346'
	// 7:11 kleisma down
	AccSagittal7v11KleismaDown = '\uE341'
	// 7:11 kleisma up, (7:11k, ~29k)
	AccSagittal7v11KleismaUp = '\uE340'
	// 7:19 comma down
	AccSagittal7v19CommaDown = '\uE39B'
	// 7:19 comma up, (7:19C, 7C less 19s)
	AccSagittal7v19CommaUp = '\uE39A'
	// 8 tinas down, 11⋅17/(5²⋅7)-schismina down, 1.14 cents down
	AccSagittal8TinasDown = '\uE407'
	// 8 tinas up, 11⋅17/(5²⋅7)-schismina up, 1.14 cents up
	AccSagittal8TinasUp = '\uE406'
	// 9 tinas down, 1/(7²⋅11)-schismina down, 1.26 cents down
	AccSagittal9TinasDown = '\uE409'
	// 9 tinas up, 1/(7²⋅11)-schismina up, 1.26 cents up
	AccSagittal9TinasUp = '\uE408'
	// Acute, 5 schisma up (5s), 2 cents up
	AccSagittalAcute = '\uE3F2'
	// Double flat, (2 apotomes down)[almost all EDOs], whole-tone down
	AccSagittalDoubleFlat = '\uE335'
	// Double flat 11:49C-up
	AccSagittalDoubleFlat11v49CUp = '\uE3E9'
	// Double flat 143C-up
	AccSagittalDoubleFlat143CUp = '\uE3EB'
	// Double flat 17C-up
	AccSagittalDoubleFlat17CUp = '\uE365'
	// Double flat 17k-up
	AccSagittalDoubleFlat17kUp = '\uE3ED'
	// Double flat 19C-up
	AccSagittalDoubleFlat19CUp = '\uE3E7'
	// Double flat 19s-up
	AccSagittalDoubleFlat19sUp = '\uE3EF'
	// Double flat 23C-up, 14° down [96 EDO], 7/8-tone down
	AccSagittalDoubleFlat23CUp = '\uE387'
	// Double flat 23S-up
	AccSagittalDoubleFlat23SUp = '\uE3E1'
	// Double flat 25S-up, 8°down [53 EDO]
	AccSagittalDoubleFlat25SUp = '\uE32D'
	// Double flat 49S-up
	AccSagittalDoubleFlat49SUp = '\uE3E3'
	// Double flat 55C-up, 13° down [96 EDO], 13/16-tone down
	AccSagittalDoubleFlat55CUp = '\uE363'
	// Double flat 5C-up, 5°[22 29] 7°[34 41] 9°53 down, 11/12 tone down
	AccSagittalDoubleFlat5CUp = '\uE331'
	// Double flat 5:11S-up
	AccSagittalDoubleFlat5v11SUp = '\uE35F'
	// Double flat 5:19C-up, 19/20-tone down
	AccSagittalDoubleFlat5v19CUp = '\uE385'
	// Double flat 5:23S-up, 8° down [60 EDO], 4/5-tone down
	AccSagittalDoubleFlat5v23SUp = '\uE383'
	// Double flat 5:7k-up
	AccSagittalDoubleFlat5v7kUp = '\uE333'
	// Double flat 7C-up, 5° down [43 EDO], 10° down [72 EDO], 5/6-tone down
	AccSagittalDoubleFlat7CUp = '\uE32F'
	// Double flat 7:11C-up, 9° down [60 EDO], 9/10-tone down
	AccSagittalDoubleFlat7v11CUp = '\uE361'
	// Double flat 7:11k-up
	AccSagittalDoubleFlat7v11kUp = '\uE367'
	// Double flat 7:19C-up
	AccSagittalDoubleFlat7v19CUp = '\uE3E5'
	// Double sharp, (2 apotomes up)[almost all EDOs], whole-tone up
	AccSagittalDoubleSharp = '\uE334'
	// Double sharp 11:49C-down
	AccSagittalDoubleSharp11v49CDown = '\uE3E8'
	// Double sharp 143C-down
	AccSagittalDoubleSharp143CDown = '\uE3EA'
	// Double sharp 17C-down
	AccSagittalDoubleSharp17CDown = '\uE364'
	// Double sharp 17k-down
	AccSagittalDoubleSharp17kDown = '\uE3EC'
	// Double sharp 19C-down
	AccSagittalDoubleSharp19CDown = '\uE3E6'
	// Double sharp 19s-down
	AccSagittalDoubleSharp19sDown = '\uE3EE'
	// Double sharp 23C-down, 14°up [96 EDO], 7/8-tone up
	AccSagittalDoubleSharp23CDown = '\uE386'
	// Double sharp 23S-down
	AccSagittalDoubleSharp23SDown = '\uE3E0'
	// Double sharp 25S-down, 8°up [53 EDO]
	AccSagittalDoubleSharp25SDown = '\uE32C'
	// Double sharp 49S-down
	AccSagittalDoubleSharp49SDown = '\uE3E2'
	// Double sharp 55C-down, 13° up [96 EDO], 13/16-tone up
	AccSagittalDoubleSharp55CDown = '\uE362'
	// Double sharp 5C-down, 5°[22 29] 7°[34 41] 9°53 up, 11/12 tone up
	AccSagittalDoubleSharp5CDown = '\uE330'
	// Double sharp 5:11S-down
	AccSagittalDoubleSharp5v11SDown = '\uE35E'
	// Double sharp 5:19C-down, 19/20-tone up
	AccSagittalDoubleSharp5v19CDown = '\uE384'
	// Double sharp 5:23S-down, 8° up [60 EDO], 4/5-tone up
	AccSagittalDoubleSharp5v23SDown = '\uE382'
	// Double sharp 5:7k-down
	AccSagittalDoubleSharp5v7kDown = '\uE332'
	// Double sharp 7C-down, 5°[43] 10°[72] up, 5/6-tone up
	AccSagittalDoubleSharp7CDown = '\uE32E'
	// Double sharp 7:11C-down, 9° up [60 EDO], 9/10-tone up
	AccSagittalDoubleSharp7v11CDown = '\uE360'
	// Double sharp 7:11k-down
	AccSagittalDoubleSharp7v11kDown = '\uE366'
	// Double sharp 7:19C-down
	AccSagittalDoubleSharp7v19CDown = '\uE3E4'
	// Flat, (apotome down)[almost all EDOs], 1/2-tone down
	AccSagittalFlat = '\uE319'
	// Flat 11L-down, 8° up [46 EDO]
	AccSagittalFlat11LDown = '\uE329'
	// Flat 11M-down, 3° down [17 31 EDOs], 7° down [46 EDO], 3/4-tone down
	AccSagittalFlat11MDown = '\uE327'
	// Flat 11:19L-down
	AccSagittalFlat11v19LDown = '\uE3DB'
	// Flat 11:19M-down
	AccSagittalFlat11v19MDown = '\uE3D3'
	// Flat 11:49C-down
	AccSagittalFlat11v49CDown = '\uE3C7'
	// Flat 11:49C-up
	AccSagittalFlat11v49CUp = '\uE3B9'
	// Flat 143C-down
	AccSagittalFlat143CDown = '\uE3C5'
	// Flat 143C-up
	AccSagittalFlat143CUp = '\uE3BB'
	// Flat 17C-down
	AccSagittalFlat17CDown = '\uE357'
	// Flat 17C-up
	AccSagittalFlat17CUp = '\uE351'
	// Flat 17k-down
	AccSagittalFlat17kDown = '\uE3C3'
	// Flat 17k-up
	AccSagittalFlat17kUp = '\uE3BD'
	// Flat 19C-down
	AccSagittalFlat19CDown = '\uE3C9'
	// Flat 19C-up
	AccSagittalFlat19CUp = '\uE3B7'
	// Flat 19s-down
	AccSagittalFlat19sDown = '\uE3C1'
	// Flat 19s-up
	AccSagittalFlat19sUp = '\uE3BF'
	// Flat 23C-down, 10° down [96 EDO], 5/8-tone down
	AccSagittalFlat23CDown = '\uE37D'
	// Flat 23C-up, 6° down [96 EDO], 3/8-tone down
	AccSagittalFlat23CUp = '\uE37B'
	// Flat 23S-down
	AccSagittalFlat23SDown = '\uE3CF'
	// Flat 23S-up
	AccSagittalFlat23SUp = '\uE3B1'
	// Flat 25S-down, 7° down [53 EDO]
	AccSagittalFlat25SDown = '\uE323'
	// Flat 25S-up, 3° down [53 EDO]
	AccSagittalFlat25SUp = '\uE311'
	// Flat 35L-down, 5° down [50 EDO]
	AccSagittalFlat35LDown = '\uE32B'
	// Flat 35M-down, 4° down [50 EDO], 6° down [27 EDO], 13/18-tone down
	AccSagittalFlat35MDown = '\uE325'
	// Flat 49L-down
	AccSagittalFlat49LDown = '\uE3D9'
	// Flat 49M-down
	AccSagittalFlat49MDown = '\uE3D5'
	// Flat 49S-down
	AccSagittalFlat49SDown = '\uE3CD'
	// Flat 49S-up
	AccSagittalFlat49SUp = '\uE3B3'
	// Flat 55C-down, 11° down [96 EDO], 11/16-tone down
	AccSagittalFlat55CDown = '\uE359'
	// Flat 55C-up, 5° down [96 EDO], 5/16-tone down
	AccSagittalFlat55CUp = '\uE34F'
	// Flat 5C-down, 4°[22 29] 5°[27 34 41] 6°[39 46 53] down, 7/12-tone down
	AccSagittalFlat5CDown = '\uE31F'
	// Flat 5C-up, 2°[22 29] 3°[27 34 41] 4°[39 46 53] 5°72 7°[96] down, 5/12-tone down
	AccSagittalFlat5CUp = '\uE315'
	// Flat 5:11S-down
	AccSagittalFlat5v11SDown = '\uE35D'
	// Flat 5:11S-up
	AccSagittalFlat5v11SUp = '\uE34B'
	// Flat 5:13L-down
	AccSagittalFlat5v13LDown = '\uE3DD'
	// Flat 5:13M-down
	AccSagittalFlat5v13MDown = '\uE3D1'
	// Flat 5:19C-down, 11/20-tone down
	AccSagittalFlat5v19CDown = '\uE37F'
	// Flat 5:19C-up, 9/20-tone down
	AccSagittalFlat5v19CUp = '\uE379'
	// Flat 5:23S-down, 7° down [60 EDO], 7/10-tone down
	AccSagittalFlat5v23SDown = '\uE381'
	// Flat 5:23S-up, 3° down [60 EDO], 3/10-tone down
	AccSagittalFlat5v23SUp = '\uE377'
	// Flat 5:49M-down
	AccSagittalFlat5v49MDown = '\uE3D7'
	// Flat 5:7k-down
	AccSagittalFlat5v7kDown = '\uE31D'
	// Flat 5:7k-up
	AccSagittalFlat5v7kUp = '\uE317'
	// Flat 7C-down, 4° down [43 EDO], 8° down [72 EDO], 2/3-tone down
	AccSagittalFlat7CDown = '\uE321'
	// Flat 7C-up, 2° down [43 EDO], 4° down [72 EDO], 1/3-tone down
	AccSagittalFlat7CUp = '\uE313'
	// Flat 7:11C-down, 6° down [60 EDO], 3/5- tone down
	AccSagittalFlat7v11CDown = '\uE35B'
	// Flat 7:11C-up, 4° down [60 EDO], 2/5-tone down
	AccSagittalFlat7v11CUp = '\uE34D'
	// Flat 7:11k-down
	AccSagittalFlat7v11kDown = '\uE355'
	// Flat 7:11k-up
	AccSagittalFlat7v11kUp = '\uE353'
	// Flat 7:19C-down
	AccSagittalFlat7v19CDown = '\uE3CB'
	// Flat 7:19C-up
	AccSagittalFlat7v19CUp = '\uE3B5'
	// Fractional tina down, 77/(5⋅37)-schismina down, 0.08 cents down
	AccSagittalFractionalTinaDown = '\uE40B'
	// Fractional tina up, 77/(5⋅37)-schismina up, 0.08 cents up
	AccSagittalFractionalTinaUp = '\uE40A'
	// Grave, 5 schisma down, 2 cents down
	AccSagittalGrave = '\uE3F3'
	// Shaft down, (natural for use with only diacritics down)
	AccSagittalShaftDown = '\uE3F1'
	// Shaft up, (natural for use with only diacritics up)
	AccSagittalShaftUp = '\uE3F0'
	// Sharp, (apotome up)[almost all EDOs], 1/2-tone up
	AccSagittalSharp = '\uE318'
	// Sharp 11L-up, 8° up [46 EDO]
	AccSagittalSharp11LUp = '\uE328'
	// Sharp 11M-up, 3° up [17 31 EDOs], 7° up [46 EDO], 3/4-tone up
	AccSagittalSharp11MUp = '\uE326'
	// Sharp 11:19L-up
	AccSagittalSharp11v19LUp = '\uE3DA'
	// Sharp 11:19M-up
	AccSagittalSharp11v19MUp = '\uE3D2'
	// Sharp 11:49C-down
	AccSagittalSharp11v49CDown = '\uE3B8'
	// Sharp 11:49C-up
	AccSagittalSharp11v49CUp = '\uE3C6'
	// Sharp 143C-down
	AccSagittalSharp143CDown = '\uE3BA'
	// Sharp 143C-up
	AccSagittalSharp143CUp = '\uE3C4'
	// Sharp 17C-down
	AccSagittalSharp17CDown = '\uE350'
	// Sharp 17C-up
	AccSagittalSharp17CUp = '\uE356'
	// Sharp 17k-down
	AccSagittalSharp17kDown = '\uE3BC'
	// Sharp 17k-up
	AccSagittalSharp17kUp = '\uE3C2'
	// Sharp 19C-down
	AccSagittalSharp19CDown = '\uE3B6'
	// Sharp 19C-up
	AccSagittalSharp19CUp = '\uE3C8'
	// Sharp 19s-down
	AccSagittalSharp19sDown = '\uE3BE'
	// Sharp 19s-up
	AccSagittalSharp19sUp = '\uE3C0'
	// Sharp 23C-down, 6° up [96 EDO], 3/8-tone up
	AccSagittalSharp23CDown = '\uE37A'
	// Sharp 23C-up, 10° up [96 EDO], 5/8-tone up
	AccSagittalSharp23CUp = '\uE37C'
	// Sharp 23S-down
	AccSagittalSharp23SDown = '\uE3B0'
	// Sharp 23S-up
	AccSagittalSharp23SUp = '\uE3CE'
	// Sharp 25S-down, 3° up [53 EDO]
	AccSagittalSharp25SDown = '\uE310'
	// Sharp 25S-up, 7° up [53 EDO]
	AccSagittalSharp25SUp = '\uE322'
	// Sharp 35L-up, 5° up [50 EDO]
	AccSagittalSharp35LUp = '\uE32A'
	// Sharp 35M-up, 4° up [50 EDO], 6° up [27 EDO], 13/18-tone up
	AccSagittalSharp35MUp = '\uE324'
	// Sharp 49L-up
	AccSagittalSharp49LUp = '\uE3D8'
	// Sharp 49M-up
	AccSagittalSharp49MUp = '\uE3D4'
	// Sharp 49S-down
	AccSagittalSharp49SDown = '\uE3B2'
	// Sharp 49S-up
	AccSagittalSharp49SUp = '\uE3CC'
	// Sharp 55C-down, 5° up [96 EDO], 5/16-tone up
	AccSagittalSharp55CDown = '\uE34E'
	// Sharp 55C-up, 11° up [96 EDO], 11/16-tone up
	AccSagittalSharp55CUp = '\uE358'
	// Sharp 5C-down, 2°[22 29] 3°[27 34 41] 4°[39 46 53] 5°[72] 7°[96] up, 5/12-tone up
	AccSagittalSharp5CDown = '\uE314'
	// Sharp 5C-up, 4°[22 29] 5°[27 34 41] 6°[39 46 53] up, 7/12-tone up
	AccSagittalSharp5CUp = '\uE31E'
	// Sharp 5:11S-down
	AccSagittalSharp5v11SDown = '\uE34A'
	// Sharp 5:11S-up
	AccSagittalSharp5v11SUp = '\uE35C'
	// Sharp 5:13L-up
	AccSagittalSharp5v13LUp = '\uE3DC'
	// Sharp 5:13M-up
	AccSagittalSharp5v13MUp = '\uE3D0'
	// Sharp 5:19C-down, 9/20-tone up
	AccSagittalSharp5v19CDown = '\uE378'
	// Sharp 5:19C-up, 11/20-tone up
	AccSagittalSharp5v19CUp = '\uE37E'
	// Sharp 5:23S-down, 3° up [60 EDO], 3/10-tone up
	AccSagittalSharp5v23SDown = '\uE376'
	// Sharp 5:23S-up, 7° up [60 EDO], 7/10-tone up
	AccSagittalSharp5v23SUp = '\uE380'
	// Sharp 5:49M-up, (one and a half apotomes)
	AccSagittalSharp5v49MUp = '\uE3D6'
	// Sharp 5:7k-down
	AccSagittalSharp5v7kDown = '\uE316'
	// Sharp 5:7k-up
	AccSagittalSharp5v7kUp = '\uE31C'
	// Sharp 7C-down, 2° up [43 EDO], 4° up [72 EDO], 1/3-tone up
	AccSagittalSharp7CDown = '\uE312'
	// Sharp 7C-up, 4° up [43 EDO], 8° up [72 EDO], 2/3-tone up
	AccSagittalSharp7CUp = '\uE320'
	// Sharp 7:11C-down, 4° up [60 EDO], 2/5-tone up
	AccSagittalSharp7v11CDown = '\uE34C'
	// Sharp 7:11C-up, 6° up [60 EDO], 3/5- tone up
	AccSagittalSharp7v11CUp = '\uE35A'
	// Sharp 7:11k-down
	AccSagittalSharp7v11kDown = '\uE352'
	// Sharp 7:11k-up
	AccSagittalSharp7v11kUp = '\uE354'
	// Sharp 7:19C-down
	AccSagittalSharp7v19CDown = '\uE3B4'
	// Sharp 7:19C-up
	AccSagittalSharp7v19CUp = '\uE3CA'
	// Unused
	AccSagittalUnused1 = '\uE31A'
	// Unused
	AccSagittalUnused2 = '\uE31B'
	// Unused
	AccSagittalUnused3 = '\uE3DE'
	// Unused
	AccSagittalUnused4 = '\uE3DF'
	// Combining accordion coupler dot
	AccdnCombDot = '\uE8CA'
	// Combining left hand, 2 ranks, empty
	AccdnCombLH2RanksEmpty = '\uE8C8'
	// Combining left hand, 3 ranks, empty (square)
	AccdnCombLH3RanksEmptySquare = '\uE8C9'
	// Combining right hand, 3 ranks, empty
	AccdnCombRH3RanksEmpty = '\uE8C6'
	// Combining right hand, 4 ranks, empty
	AccdnCombRH4RanksEmpty = '\uE8C7'
	// Diatonic accordion clef
	AccdnDiatonicClef = '\uE079'
	// Left hand, 2 ranks, 16' stop (round)
	AccdnLH2Ranks16Round = '\uE8BC'
	// Left hand, 2 ranks, 8' stop + 16' stop (round)
	AccdnLH2Ranks8Plus16Round = '\uE8BD'
	// Left hand, 2 ranks, 8' stop (round)
	AccdnLH2Ranks8Round = '\uE8BB'
	// Left hand, 2 ranks, full master (round)
	AccdnLH2RanksFullMasterRound = '\uE8C0'
	// Left hand, 2 ranks, master + 16' stop (round)
	AccdnLH2RanksMasterPlus16Round = '\uE8BF'
	// Left hand, 2 ranks, master (round)
	AccdnLH2RanksMasterRound = '\uE8BE'
	// Left hand, 3 ranks, 2' stop + 8' stop (square)
	AccdnLH3Ranks2Plus8Square = '\uE8C4'
	// Left hand, 3 ranks, 2' stop (square)
	AccdnLH3Ranks2Square = '\uE8C2'
	// Left hand, 3 ranks, 8' stop (square)
	AccdnLH3Ranks8Square = '\uE8C1'
	// Left hand, 3 ranks, double 8' stop (square)
	AccdnLH3RanksDouble8Square = '\uE8C3'
	// Left hand, 3 ranks, 2' stop + double 8' stop (tutti) (square)
	AccdnLH3RanksTuttiSquare = '\uE8C5'
	// Pull
	AccdnPull = '\uE8CC'
	// Push
	AccdnPush = '\uE8CB'
	// Right hand, 3 ranks, 8' stop + upper tremolo 8' stop + 16' stop (accordion)
	AccdnRH3RanksAccordion = '\uE8AC'
	// Right hand, 3 ranks, lower tremolo 8' stop + 8' stop + upper tremolo 8' stop (authentic musette)
	AccdnRH3RanksAuthenticMusette = '\uE8A8'
	// Right hand, 3 ranks, 8' stop + 16' stop (bandoneón)
	AccdnRH3RanksBandoneon = '\uE8AB'
	// Right hand, 3 ranks, 16' stop (bassoon)
	AccdnRH3RanksBassoon = '\uE8A4'
	// Right hand, 3 ranks, 8' stop (clarinet)
	AccdnRH3RanksClarinet = '\uE8A1'
	// Right hand, 3 ranks, lower tremolo 8' stop + 8' stop + upper tremolo 8' stop + 16' stop
	AccdnRH3RanksDoubleTremoloLower8ve = '\uE8B1'
	// Right hand, 3 ranks, 4' stop + lower tremolo 8' stop + 8' stop + upper tremolo 8' stop
	AccdnRH3RanksDoubleTremoloUpper8ve = '\uE8B2'
	// Right hand, 3 ranks, 4' stop + lower tremolo 8' stop + 8' stop + upper tremolo 8' stop + 16' stop
	AccdnRH3RanksFullFactory = '\uE8B3'
	// Right hand, 3 ranks, 4' stop + 8' stop + 16' stop (harmonium)
	AccdnRH3RanksHarmonium = '\uE8AA'
	// Right hand, 3 ranks, 4' stop + 8' stop + upper tremolo 8' stop (imitation musette)
	AccdnRH3RanksImitationMusette = '\uE8A7'
	// Right hand, 3 ranks, lower tremolo 8' stop
	AccdnRH3RanksLowerTremolo8 = '\uE8A3'
	// Right hand, 3 ranks, 4' stop + lower tremolo 8' stop + upper tremolo 8' stop + 16' stop (master)
	AccdnRH3RanksMaster = '\uE8AD'
	// Right hand, 3 ranks, 4' stop + 8' stop (oboe)
	AccdnRH3RanksOboe = '\uE8A5'
	// Right hand, 3 ranks, 4' stop + 16' stop (organ)
	AccdnRH3RanksOrgan = '\uE8A9'
	// Right hand, 3 ranks, 4' stop (piccolo)
	AccdnRH3RanksPiccolo = '\uE8A0'
	// Right hand, 3 ranks, lower tremolo 8' stop + upper tremolo 8' stop + 16' stop
	AccdnRH3RanksTremoloLower8ve = '\uE8AF'
	// Right hand, 3 ranks, 4' stop + lower tremolo 8' stop + upper tremolo 8' stop
	AccdnRH3RanksTremoloUpper8ve = '\uE8B0'
	// Right hand, 3 ranks, lower tremolo 8' stop + upper tremolo 8' stop
	AccdnRH3RanksTwoChoirs = '\uE8AE'
	// Right hand, 3 ranks, upper tremolo 8' stop
	AccdnRH3RanksUpperTremolo8 = '\uE8A2'
	// Right hand, 3 ranks, 8' stop + upper tremolo 8' stop (violin)
	AccdnRH3RanksViolin = '\uE8A6'
	// Right hand, 4 ranks, alto
	AccdnRH4RanksAlto = '\uE8B5'
	// Right hand, 4 ranks, bass/alto
	AccdnRH4RanksBassAlto = '\uE8BA'
	// Right hand, 4 ranks, master
	AccdnRH4RanksMaster = '\uE8B7'
	// Right hand, 4 ranks, soft bass
	AccdnRH4RanksSoftBass = '\uE8B8'
	// Right hand, 4 ranks, soft tenor
	AccdnRH4RanksSoftTenor = '\uE8B9'
	// Right hand, 4 ranks, soprano
	AccdnRH4RanksSoprano = '\uE8B4'
	// Right hand, 4 ranks, tenor
	AccdnRH4RanksTenor = '\uE8B6'
	// Ricochet (2 tones)
	AccdnRicochet2 = '\uE8CD'
	// Ricochet (3 tones)
	AccdnRicochet3 = '\uE8CE'
	// Ricochet (4 tones)
	AccdnRicochet4 = '\uE8CF'
	// Ricochet (5 tones)
	AccdnRicochet5 = '\uE8D0'
	// Ricochet (6 tones)
	AccdnRicochet6 = '\uE8D1'
	// Combining ricochet for stem (2 tones)
	AccdnRicochetStem2 = '\uE8D2'
	// Combining ricochet for stem (3 tones)
	AccdnRicochetStem3 = '\uE8D3'
	// Combining ricochet for stem (4 tones)
	AccdnRicochetStem4 = '\uE8D4'
	// Combining ricochet for stem (5 tones)
	AccdnRicochetStem5 = '\uE8D5'
	// Combining ricochet for stem (6 tones)
	AccdnRicochetStem6 = '\uE8D6'
	// 1-comma flat
	Accidental1CommaFlat = '\uE454'
	// 1-comma sharp
	Accidental1CommaSharp = '\uE450'
	// 2-comma flat
	Accidental2CommaFlat = '\uE455'
	// 2-comma sharp
	Accidental2CommaSharp = '\uE451'
	// 3-comma flat
	Accidental3CommaFlat = '\uE456'
	// 3-comma sharp
	Accidental3CommaSharp = '\uE452'
	// 4-comma flat
	Accidental4CommaFlat = '\uE457'
	// 5-comma sharp
	Accidental5CommaSharp = '\uE453'
	// Arrow down (lower by one quarter-tone)
	AccidentalArrowDown = '\uE27B'
	// Arrow up (raise by one quarter-tone)
	AccidentalArrowUp = '\uE27A'
	// Bakiye (flat)
	AccidentalBakiyeFlat = '\uE442'
	// Bakiye (sharp)
	AccidentalBakiyeSharp = '\uE445'
	// Accidental bracket, left
	AccidentalBracketLeft = '\uE26C'
	// Accidental bracket, right
	AccidentalBracketRight = '\uE26D'
	// Büyük mücenneb (flat)
	AccidentalBuyukMucennebFlat = '\uE440'
	// Büyük mücenneb (sharp)
	AccidentalBuyukMucennebSharp = '\uE447'
	// Combining close curly brace
	AccidentalCombiningCloseCurlyBrace = '\uE2EF'
	// Combining lower by one 17-limit schisma
	AccidentalCombiningLower17Schisma = '\uE2E6'
	// Combining lower by one 19-limit schisma
	AccidentalCombiningLower19Schisma = '\uE2E8'
	// Combining lower by one 23-limit comma
	AccidentalCombiningLower23Limit29LimitComma = '\uE2EA'
	// Combining lower by one 29-limit comma
	AccidentalCombiningLower29LimitComma = '\uEE50'
	// Combining lower by one 31-limit schisma
	AccidentalCombiningLower31Schisma = '\uE2EC'
	// Combining lower by one 37-limit quartertone
	AccidentalCombiningLower37Quartertone = '\uEE52'
	// Combining lower by one 41-limit comma
	AccidentalCombiningLower41Comma = '\uEE54'
	// Combining lower by one 43-limit comma
	AccidentalCombiningLower43Comma = '\uEE56'
	// Combining lower by one 47-limit quartertone
	AccidentalCombiningLower47Quartertone = '\uEE58'
	// Combining lower by one 53-limit comma
	AccidentalCombiningLower53LimitComma = '\uE2F7'
	// Combining open curly brace
	AccidentalCombiningOpenCurlyBrace = '\uE2EE'
	// Combining raise by one 17-limit schisma
	AccidentalCombiningRaise17Schisma = '\uE2E7'
	// Combining raise by one 19-limit schisma
	AccidentalCombiningRaise19Schisma = '\uE2E9'
	// Combining raise by one 23-limit comma
	AccidentalCombiningRaise23Limit29LimitComma = '\uE2EB'
	// Combining raise by one 29-limit comma
	AccidentalCombiningRaise29LimitComma = '\uEE51'
	// Combining raise by one 31-limit schisma
	AccidentalCombiningRaise31Schisma = '\uE2ED'
	// Combining raise by one 37-limit quartertone
	AccidentalCombiningRaise37Quartertone = '\uEE53'
	// Combining raise by one 41-limit comma
	AccidentalCombiningRaise41Comma = '\uEE55'
	// Combining raise by one 43-limit comma
	AccidentalCombiningRaise43Comma = '\uEE57'
	// Combining raise by one 47-limit quartertone
	AccidentalCombiningRaise47Quartertone = '\uEE59'
	// Combining raise by one 53-limit comma
	AccidentalCombiningRaise53LimitComma = '\uE2F8'
	// Syntonic/Didymus comma (80:81) down (Bosanquet)
	AccidentalCommaSlashDown = '\uE47A'
	// Syntonic/Didymus comma (80:81) up (Bosanquet)
	AccidentalCommaSlashUp = '\uE479'
	// Double flat
	AccidentalDoubleFlat = '\uE264'
	// Arabic double flat
	AccidentalDoubleFlatArabic = '\uED30'
	// Double flat equal tempered semitone
	AccidentalDoubleFlatEqualTempered = '\uE2F0'
	// Double flat lowered by one syntonic comma
	AccidentalDoubleFlatOneArrowDown = '\uE2C0'
	// Double flat raised by one syntonic comma
	AccidentalDoubleFlatOneArrowUp = '\uE2C5'
	// Reversed double flat
	AccidentalDoubleFlatReversed = '\uE483'
	// Double flat lowered by three syntonic commas
	AccidentalDoubleFlatThreeArrowsDown = '\uE2D4'
	// Double flat raised by three syntonic commas
	AccidentalDoubleFlatThreeArrowsUp = '\uE2D9'
	// Turned double flat
	AccidentalDoubleFlatTurned = '\uE485'
	// Double flat lowered by two syntonic commas
	AccidentalDoubleFlatTwoArrowsDown = '\uE2CA'
	// Double flat raised by two syntonic commas
	AccidentalDoubleFlatTwoArrowsUp = '\uE2CF'
	// Double sharp
	AccidentalDoubleSharp = '\uE263'
	// Arabic double sharp
	AccidentalDoubleSharpArabic = '\uED38'
	// Double sharp equal tempered semitone
	AccidentalDoubleSharpEqualTempered = '\uE2F4'
	// Double sharp lowered by one syntonic comma
	AccidentalDoubleSharpOneArrowDown = '\uE2C4'
	// Double sharp raised by one syntonic comma
	AccidentalDoubleSharpOneArrowUp = '\uE2C9'
	// Double sharp lowered by three syntonic commas
	AccidentalDoubleSharpThreeArrowsDown = '\uE2D8'
	// Double sharp raised by three syntonic commas
	AccidentalDoubleSharpThreeArrowsUp = '\uE2DD'
	// Double sharp lowered by two syntonic commas
	AccidentalDoubleSharpTwoArrowsDown = '\uE2CE'
	// Double sharp raised by two syntonic commas
	AccidentalDoubleSharpTwoArrowsUp = '\uE2D3'
	// Enharmonically reinterpret accidental almost equal to
	AccidentalEnharmonicAlmostEqualTo = '\uE2FA'
	// Enharmonically reinterpret accidental equals
	AccidentalEnharmonicEquals = '\uE2FB'
	// Enharmonically reinterpret accidental tilde
	AccidentalEnharmonicTilde = '\uE2F9'
	// Filled reversed flat and flat
	AccidentalFilledReversedFlatAndFlat = '\uE296'
	// Filled reversed flat and flat with arrow down
	AccidentalFilledReversedFlatAndFlatArrowDown = '\uE298'
	// Filled reversed flat and flat with arrow up
	AccidentalFilledReversedFlatAndFlatArrowUp = '\uE297'
	// Filled reversed flat with arrow down
	AccidentalFilledReversedFlatArrowDown = '\uE293'
	// Filled reversed flat with arrow up
	AccidentalFilledReversedFlatArrowUp = '\uE292'
	// Five-quarter-tones flat
	AccidentalFiveQuarterTonesFlatArrowDown = '\uE279'
	// Five-quarter-tones sharp
	AccidentalFiveQuarterTonesSharpArrowUp = '\uE276'
	// Flat
	AccidentalFlat = '\uE260'
	// Arabic half-tone flat
	AccidentalFlatArabic = '\uED32'
	// Flat equal tempered semitone
	AccidentalFlatEqualTempered = '\uE2F1'
	// Lowered flat (Stockhausen)
	AccidentalFlatLoweredStockhausen = '\uED53'
	// Flat lowered by one syntonic comma
	AccidentalFlatOneArrowDown = '\uE2C1'
	// Flat raised by one syntonic comma
	AccidentalFlatOneArrowUp = '\uE2C6'
	// Raised flat (Stockhausen)
	AccidentalFlatRaisedStockhausen = '\uED52'
	// Repeated flat, note on line (Stockhausen)
	AccidentalFlatRepeatedLineStockhausen = '\uED5C'
	// Repeated flat, note in space (Stockhausen)
	AccidentalFlatRepeatedSpaceStockhausen = '\uED5B'
	// Flat lowered by three syntonic commas
	AccidentalFlatThreeArrowsDown = '\uE2D5'
	// Flat raised by three syntonic commas
	AccidentalFlatThreeArrowsUp = '\uE2DA'
	// Turned flat
	AccidentalFlatTurned = '\uE484'
	// Flat lowered by two syntonic commas
	AccidentalFlatTwoArrowsDown = '\uE2CB'
	// Flat raised by two syntonic commas
	AccidentalFlatTwoArrowsUp = '\uE2D0'
	// Quarter-tone higher (Alois Hába)
	AccidentalHabaFlatQuarterToneHigher = '\uEE65'
	// Three quarter-tones lower (Alois Hába)
	AccidentalHabaFlatThreeQuarterTonesLower = '\uEE69'
	// Quarter-tone higher (Alois Hába)
	AccidentalHabaQuarterToneHigher = '\uEE64'
	// Quarter-tone lower (Alois Hába)
	AccidentalHabaQuarterToneLower = '\uEE67'
	// Quarter-tone lower (Alois Hába)
	AccidentalHabaSharpQuarterToneLower = '\uEE68'
	// Three quarter-tones higher (Alois Hába)
	AccidentalHabaSharpThreeQuarterTonesHigher = '\uEE66'
	// Half sharp with arrow down
	AccidentalHalfSharpArrowDown = '\uE29A'
	// Half sharp with arrow up
	AccidentalHalfSharpArrowUp = '\uE299'
	// Thirteen (raise by 65:64)
	AccidentalJohnston13 = '\uE2B6'
	// Inverted 13 (lower by 65:64)
	AccidentalJohnston31 = '\uE2B7'
	// Down arrow (lower by 33:32)
	AccidentalJohnstonDown = '\uE2B5'
	// Inverted seven (raise by 36:35)
	AccidentalJohnstonEl = '\uE2B2'
	// Minus (lower by 81:80)
	AccidentalJohnstonMinus = '\uE2B1'
	// Plus (raise by 81:80)
	AccidentalJohnstonPlus = '\uE2B0'
	// Seven (lower by 36:35)
	AccidentalJohnstonSeven = '\uE2B3'
	// Up arrow (raise by 33:32)
	AccidentalJohnstonUp = '\uE2B4'
	// Koma (flat)
	AccidentalKomaFlat = '\uE443'
	// Koma (sharp)
	AccidentalKomaSharp = '\uE444'
	// Koron (quarter tone flat)
	AccidentalKoron = '\uE460'
	// Küçük mücenneb (flat)
	AccidentalKucukMucennebFlat = '\uE441'
	// Küçük mücenneb (sharp)
	AccidentalKucukMucennebSharp = '\uE446'
	// Large double sharp
	AccidentalLargeDoubleSharp = '\uE47D'
	// Lower by one septimal comma
	AccidentalLowerOneSeptimalComma = '\uE2DE'
	// Lower by one tridecimal quartertone
	AccidentalLowerOneTridecimalQuartertone = '\uE2E4'
	// Lower by one undecimal quartertone
	AccidentalLowerOneUndecimalQuartertone = '\uE2E2'
	// Lower by two septimal commas
	AccidentalLowerTwoSeptimalCommas = '\uE2E0'
	// Lowered (Stockhausen)
	AccidentalLoweredStockhausen = '\uED51'
	// Narrow reversed flat(quarter-tone flat)
	AccidentalNarrowReversedFlat = '\uE284'
	// Narrow reversed flat and flat(three-quarter-tones flat)
	AccidentalNarrowReversedFlatAndFlat = '\uE285'
	// Natural
	AccidentalNatural = '\uE261'
	// Arabic natural
	AccidentalNaturalArabic = '\uED34'
	// Natural equal tempered semitone
	AccidentalNaturalEqualTempered = '\uE2F2'
	// Natural flat
	AccidentalNaturalFlat = '\uE267'
	// Lowered natural (Stockhausen)
	AccidentalNaturalLoweredStockhausen = '\uED55'
	// Natural lowered by one syntonic comma
	AccidentalNaturalOneArrowDown = '\uE2C2'
	// Natural raised by one syntonic comma
	AccidentalNaturalOneArrowUp = '\uE2C7'
	// Raised natural (Stockhausen)
	AccidentalNaturalRaisedStockhausen = '\uED54'
	// Reversed natural
	AccidentalNaturalReversed = '\uE482'
	// Natural sharp
	AccidentalNaturalSharp = '\uE268'
	// Natural lowered by three syntonic commas
	AccidentalNaturalThreeArrowsDown = '\uE2D6'
	// Natural raised by three syntonic commas
	AccidentalNaturalThreeArrowsUp = '\uE2DB'
	// Natural lowered by two syntonic commas
	AccidentalNaturalTwoArrowsDown = '\uE2CC'
	// Natural raised by two syntonic commas
	AccidentalNaturalTwoArrowsUp = '\uE2D1'
	// One and a half sharps with arrow down
	AccidentalOneAndAHalfSharpsArrowDown = '\uE29C'
	// One and a half sharps with arrow up
	AccidentalOneAndAHalfSharpsArrowUp = '\uE29B'
	// One-quarter-tone flat (Ferneyhough)
	AccidentalOneQuarterToneFlatFerneyhough = '\uE48F'
	// One-quarter-tone flat (Stockhausen)
	AccidentalOneQuarterToneFlatStockhausen = '\uED59'
	// One-quarter-tone sharp (Ferneyhough)
	AccidentalOneQuarterToneSharpFerneyhough = '\uE48E'
	// One-quarter-tone sharp (Stockhausen)
	AccidentalOneQuarterToneSharpStockhausen = '\uED58'
	// One-third-tone flat (Ferneyhough)
	AccidentalOneThirdToneFlatFerneyhough = '\uE48B'
	// One-third-tone sharp (Ferneyhough)
	AccidentalOneThirdToneSharpFerneyhough = '\uE48A'
	// Accidental parenthesis, left
	AccidentalParensLeft = '\uE26A'
	// Accidental parenthesis, right
	AccidentalParensRight = '\uE26B'
	// Lower by one equal tempered quarter-tone
	AccidentalQuarterFlatEqualTempered = '\uE2F5'
	// Raise by one equal tempered quarter tone
	AccidentalQuarterSharpEqualTempered = '\uE2F6'
	// Quarter-tone flat
	AccidentalQuarterToneFlat4 = '\uE47F'
	// Arabic quarter-tone flat
	AccidentalQuarterToneFlatArabic = '\uED33'
	// Quarter-tone flat
	AccidentalQuarterToneFlatArrowUp = '\uE270'
	// Filled reversed flat (quarter-tone flat)
	AccidentalQuarterToneFlatFilledReversed = '\uE480'
	// Quarter-tone flat
	AccidentalQuarterToneFlatNaturalArrowDown = '\uE273'
	// Quarter tone flat (Penderecki)
	AccidentalQuarterToneFlatPenderecki = '\uE478'
	// Reversed flat (quarter-tone flat) (Stein)
	AccidentalQuarterToneFlatStein = '\uE280'
	// Quarter-tone flat (van Blankenburg)
	AccidentalQuarterToneFlatVanBlankenburg = '\uE488'
	// Quarter-tone sharp
	AccidentalQuarterToneSharp4 = '\uE47E'
	// Arabic quarter-tone sharp
	AccidentalQuarterToneSharpArabic = '\uED35'
	// Quarter-tone sharp
	AccidentalQuarterToneSharpArrowDown = '\uE275'
	// Quarter tone sharp (Bussotti)
	AccidentalQuarterToneSharpBusotti = '\uE472'
	// Quarter-tone sharp
	AccidentalQuarterToneSharpNaturalArrowUp = '\uE272'
	// Half sharp (quarter-tone sharp) (Stein)
	AccidentalQuarterToneSharpStein = '\uE282'
	// Quarter tone sharp with wiggly tail
	AccidentalQuarterToneSharpWiggle = '\uE475'
	// Raise by one septimal comma
	AccidentalRaiseOneSeptimalComma = '\uE2DF'
	// Raise by one tridecimal quartertone
	AccidentalRaiseOneTridecimalQuartertone = '\uE2E5'
	// Raise by one undecimal quartertone
	AccidentalRaiseOneUndecimalQuartertone = '\uE2E3'
	// Raise by two septimal commas
	AccidentalRaiseTwoSeptimalCommas = '\uE2E1'
	// Raised (Stockhausen)
	AccidentalRaisedStockhausen = '\uED50'
	// Reversed flat and flat with arrow down
	AccidentalReversedFlatAndFlatArrowDown = '\uE295'
	// Reversed flat and flat with arrow up
	AccidentalReversedFlatAndFlatArrowUp = '\uE294'
	// Reversed flat with arrow down
	AccidentalReversedFlatArrowDown = '\uE291'
	// Reversed flat with arrow up
	AccidentalReversedFlatArrowUp = '\uE290'
	// Sharp
	AccidentalSharp = '\uE262'
	// Arabic half-tone sharp
	AccidentalSharpArabic = '\uED36'
	// Sharp equal tempered semitone
	AccidentalSharpEqualTempered = '\uE2F3'
	// Lowered sharp (Stockhausen)
	AccidentalSharpLoweredStockhausen = '\uED57'
	// Sharp lowered by one syntonic comma
	AccidentalSharpOneArrowDown = '\uE2C3'
	// Sharp raised by one syntonic comma
	AccidentalSharpOneArrowUp = '\uE2C8'
	// One or three quarter tones sharp
	AccidentalSharpOneHorizontalStroke = '\uE473'
	// Raised sharp (Stockhausen)
	AccidentalSharpRaisedStockhausen = '\uED56'
	// Repeated sharp, note on line (Stockhausen)
	AccidentalSharpRepeatedLineStockhausen = '\uED5E'
	// Repeated sharp, note in space (Stockhausen)
	AccidentalSharpRepeatedSpaceStockhausen = '\uED5D'
	// Reversed sharp
	AccidentalSharpReversed = '\uE481'
	// Sharp sharp
	AccidentalSharpSharp = '\uE269'
	// Sharp lowered by three syntonic commas
	AccidentalSharpThreeArrowsDown = '\uE2D7'
	// Sharp raised by three syntonic commas
	AccidentalSharpThreeArrowsUp = '\uE2DC'
	// Sharp lowered by two syntonic commas
	AccidentalSharpTwoArrowsDown = '\uE2CD'
	// Sharp raised by two syntonic commas
	AccidentalSharpTwoArrowsUp = '\uE2D2'
	// 1/12 tone low
	AccidentalSims12Down = '\uE2A0'
	// 1/12 tone high
	AccidentalSims12Up = '\uE2A3'
	// 1/4 tone low
	AccidentalSims4Down = '\uE2A2'
	// 1/4 tone high
	AccidentalSims4Up = '\uE2A5'
	// 1/6 tone low
	AccidentalSims6Down = '\uE2A1'
	// 1/6 tone high
	AccidentalSims6Up = '\uE2A4'
	// Sori (quarter tone sharp)
	AccidentalSori = '\uE461'
	// Byzantine-style Bakiye flat (Tavener)
	AccidentalTavenerFlat = '\uE477'
	// Byzantine-style Büyük mücenneb sharp (Tavener)
	AccidentalTavenerSharp = '\uE476'
	// Arabic three-quarter-tones flat
	AccidentalThreeQuarterTonesFlatArabic = '\uED31'
	// Three-quarter-tones flat
	AccidentalThreeQuarterTonesFlatArrowDown = '\uE271'
	// Three-quarter-tones flat
	AccidentalThreeQuarterTonesFlatArrowUp = '\uE278'
	// Three-quarter-tones flat (Couper)
	AccidentalThreeQuarterTonesFlatCouper = '\uE489'
	// Three-quarter-tones flat (Grisey)
	AccidentalThreeQuarterTonesFlatGrisey = '\uE486'
	// Three-quarter-tones flat (Tartini)
	AccidentalThreeQuarterTonesFlatTartini = '\uE487'
	// Reversed flat and flat (three-quarter-tones flat) (Zimmermann)
	AccidentalThreeQuarterTonesFlatZimmermann = '\uE281'
	// Arabic three-quarter-tones sharp
	AccidentalThreeQuarterTonesSharpArabic = '\uED37'
	// Three-quarter-tones sharp
	AccidentalThreeQuarterTonesSharpArrowDown = '\uE277'
	// Three-quarter-tones sharp
	AccidentalThreeQuarterTonesSharpArrowUp = '\uE274'
	// Three quarter tones sharp (Bussotti)
	AccidentalThreeQuarterTonesSharpBusotti = '\uE474'
	// One and a half sharps (three-quarter-tones sharp) (Stein)
	AccidentalThreeQuarterTonesSharpStein = '\uE283'
	// Three-quarter-tones sharp (Stockhausen)
	AccidentalThreeQuarterTonesSharpStockhausen = '\uED5A'
	// Triple flat
	AccidentalTripleFlat = '\uE266'
	// Triple sharp
	AccidentalTripleSharp = '\uE265'
	// Two-third-tones flat (Ferneyhough)
	AccidentalTwoThirdTonesFlatFerneyhough = '\uE48D'
	// Two-third-tones sharp (Ferneyhough)
	AccidentalTwoThirdTonesSharpFerneyhough = '\uE48C'
	// Accidental down
	AccidentalUpsAndDownsDown = '\uEE61'
	// Accidental less
	AccidentalUpsAndDownsLess = '\uEE63'
	// Accidental more
	AccidentalUpsAndDownsMore = '\uEE62'
	// Accidental up
	AccidentalUpsAndDownsUp = '\uEE60'
	// Wilson minus (5 comma down)
	AccidentalWilsonMinus = '\uE47C'
	// Wilson plus (5 comma up)
	AccidentalWilsonPlus = '\uE47B'
	// 5/6 tone flat
	AccidentalWyschnegradsky10TwelfthsFlat = '\uE434'
	// 5/6 tone sharp
	AccidentalWyschnegradsky10TwelfthsSharp = '\uE429'
	// 11/12 tone flat
	AccidentalWyschnegradsky11TwelfthsFlat = '\uE435'
	// 11/12 tone sharp
	AccidentalWyschnegradsky11TwelfthsSharp = '\uE42A'
	// 1/12 tone flat
	AccidentalWyschnegradsky1TwelfthsFlat = '\uE42B'
	// 1/12 tone sharp
	AccidentalWyschnegradsky1TwelfthsSharp = '\uE420'
	// 1/6 tone flat
	AccidentalWyschnegradsky2TwelfthsFlat = '\uE42C'
	// 1/6 tone sharp
	AccidentalWyschnegradsky2TwelfthsSharp = '\uE421'
	// 1/4 tone flat
	AccidentalWyschnegradsky3TwelfthsFlat = '\uE42D'
	// 1/4 tone sharp
	AccidentalWyschnegradsky3TwelfthsSharp = '\uE422'
	// 1/3 tone flat
	AccidentalWyschnegradsky4TwelfthsFlat = '\uE42E'
	// 1/3 tone sharp
	AccidentalWyschnegradsky4TwelfthsSharp = '\uE423'
	// 5/12 tone flat
	AccidentalWyschnegradsky5TwelfthsFlat = '\uE42F'
	// 5/12 tone sharp
	AccidentalWyschnegradsky5TwelfthsSharp = '\uE424'
	// 1/2 tone flat
	AccidentalWyschnegradsky6TwelfthsFlat = '\uE430'
	// 1/2 tone sharp
	AccidentalWyschnegradsky6TwelfthsSharp = '\uE425'
	// 7/12 tone flat
	AccidentalWyschnegradsky7TwelfthsFlat = '\uE431'
	// 7/12 tone sharp
	AccidentalWyschnegradsky7TwelfthsSharp = '\uE426'
	// 2/3 tone flat
	AccidentalWyschnegradsky8TwelfthsFlat = '\uE432'
	// 2/3 tone sharp
	AccidentalWyschnegradsky8TwelfthsSharp = '\uE427'
	// 3/4 tone flat
	AccidentalWyschnegradsky9TwelfthsFlat = '\uE433'
	// 3/4 tone sharp
	AccidentalWyschnegradsky9TwelfthsSharp = '\uE428'
	// One-third-tone sharp (Xenakis)
	AccidentalXenakisOneThirdToneSharp = '\uE470'
	// Two-third-tones sharp (Xenakis)
	AccidentalXenakisTwoThirdTonesSharp = '\uE471'
	// Choralmelodie (Berg)
	AnalyticsChoralmelodie = '\uE86A'
	// End of stimme
	AnalyticsEndStimme = '\uE863'
	// Hauptrhythmus (Berg)
	AnalyticsHauptrhythmus = '\uE86B'
	// Hauptstimme
	AnalyticsHauptstimme = '\uE860'
	// Inversion 1
	AnalyticsInversion1 = '\uE869'
	// Nebenstimme
	AnalyticsNebenstimme = '\uE861'
	// Start of stimme
	AnalyticsStartStimme = '\uE862'
	// Theme
	AnalyticsTheme = '\uE864'
	// Theme 1
	AnalyticsTheme1 = '\uE868'
	// Inversion of theme
	AnalyticsThemeInversion = '\uE867'
	// Retrograde of theme
	AnalyticsThemeRetrograde = '\uE865'
	// Retrograde inversion of theme
	AnalyticsThemeRetrogradeInversion = '\uE866'
	// Arpeggiato
	Arpeggiato = '\uE63C'
	// Arpeggiato down
	ArpeggiatoDown = '\uE635'
	// Arpeggiato up
	ArpeggiatoUp = '\uE634'
	// Black arrow down (S)
	ArrowBlackDown = '\uEB64'
	// Black arrow down-left (SW)
	ArrowBlackDownLeft = '\uEB65'
	// Black arrow down-right (SE)
	ArrowBlackDownRight = '\uEB63'
	// Black arrow left (W)
	ArrowBlackLeft = '\uEB66'
	// Black arrow right (E)
	ArrowBlackRight = '\uEB62'
	// Black arrow up (N)
	ArrowBlackUp = '\uEB60'
	// Black arrow up-left (NW)
	ArrowBlackUpLeft = '\uEB67'
	// Black arrow up-right (NE)
	ArrowBlackUpRight = '\uEB61'
	// Open arrow down (S)
	ArrowOpenDown = '\uEB74'
	// Open arrow down-left (SW)
	ArrowOpenDownLeft = '\uEB75'
	// Open arrow down-right (SE)
	ArrowOpenDownRight = '\uEB73'
	// Open arrow left (W)
	ArrowOpenLeft = '\uEB76'
	// Open arrow right (E)
	ArrowOpenRight = '\uEB72'
	// Open arrow up (N)
	ArrowOpenUp = '\uEB70'
	// Open arrow up-left (NW)
	ArrowOpenUpLeft = '\uEB77'
	// Open arrow up-right (NE)
	ArrowOpenUpRight = '\uEB71'
	// White arrow down (S)
	ArrowWhiteDown = '\uEB6C'
	// White arrow down-left (SW)
	ArrowWhiteDownLeft = '\uEB6D'
	// White arrow down-right (SE)
	ArrowWhiteDownRight = '\uEB6B'
	// White arrow left (W)
	ArrowWhiteLeft = '\uEB6E'
	// White arrow right (E)
	ArrowWhiteRight = '\uEB6A'
	// White arrow up (N)
	ArrowWhiteUp = '\uEB68'
	// White arrow up-left (NW)
	ArrowWhiteUpLeft = '\uEB6F'
	// White arrow up-right (NE)
	ArrowWhiteUpRight = '\uEB69'
	// Black arrowhead down (S)
	ArrowheadBlackDown = '\uEB7C'
	// Black arrowhead down-left (SW)
	ArrowheadBlackDownLeft = '\uEB7D'
	// Black arrowhead down-right (SE)
	ArrowheadBlackDownRight = '\uEB7B'
	// Black arrowhead left (W)
	ArrowheadBlackLeft = '\uEB7E'
	// Black arrowhead right (E)
	ArrowheadBlackRight = '\uEB7A'
	// Black arrowhead up (N)
	ArrowheadBlackUp = '\uEB78'
	// Black arrowhead up-left (NW)
	ArrowheadBlackUpLeft = '\uEB7F'
	// Black arrowhead up-right (NE)
	ArrowheadBlackUpRight = '\uEB79'
	// Open arrowhead down (S)
	ArrowheadOpenDown = '\uEB8C'
	// Open arrowhead down-left (SW)
	ArrowheadOpenDownLeft = '\uEB8D'
	// Open arrowhead down-right (SE)
	ArrowheadOpenDownRight = '\uEB8B'
	// Open arrowhead left (W)
	ArrowheadOpenLeft = '\uEB8E'
	// Open arrowhead right (E)
	ArrowheadOpenRight = '\uEB8A'
	// Open arrowhead up (N)
	ArrowheadOpenUp = '\uEB88'
	// Open arrowhead up-left (NW)
	ArrowheadOpenUpLeft = '\uEB8F'
	// Open arrowhead up-right (NE)
	ArrowheadOpenUpRight = '\uEB89'
	// White arrowhead down (S)
	ArrowheadWhiteDown = '\uEB84'
	// White arrowhead down-left (SW)
	ArrowheadWhiteDownLeft = '\uEB85'
	// White arrowhead down-right (SE)
	ArrowheadWhiteDownRight = '\uEB83'
	// White arrowhead left (W)
	ArrowheadWhiteLeft = '\uEB86'
	// White arrowhead right (E)
	ArrowheadWhiteRight = '\uEB82'
	// White arrowhead up (N)
	ArrowheadWhiteUp = '\uEB80'
	// White arrowhead up-left (NW)
	ArrowheadWhiteUpLeft = '\uEB87'
	// White arrowhead up-right (NE)
	ArrowheadWhiteUpRight = '\uEB81'
	// Accent above
	ArticAccentAbove = '\uE4A0'
	// Accent below
	ArticAccentBelow = '\uE4A1'
	// Accent-staccato above
	ArticAccentStaccatoAbove = '\uE4B0'
	// Accent-staccato below
	ArticAccentStaccatoBelow = '\uE4B1'
	// Laissez vibrer (l.v.) above
	ArticLaissezVibrerAbove = '\uE4BA'
	// Laissez vibrer (l.v.) below
	ArticLaissezVibrerBelow = '\uE4BB'
	// Marcato above
	ArticMarcatoAbove = '\uE4AC'
	// Marcato below
	ArticMarcatoBelow = '\uE4AD'
	// Marcato-staccato above
	ArticMarcatoStaccatoAbove = '\uE4AE'
	// Marcato-staccato below
	ArticMarcatoStaccatoBelow = '\uE4AF'
	// Marcato-tenuto above
	ArticMarcatoTenutoAbove = '\uE4BC'
	// Marcato-tenuto below
	ArticMarcatoTenutoBelow = '\uE4BD'
	// Soft accent above
	ArticSoftAccentAbove = '\uED40'
	// Soft accent below
	ArticSoftAccentBelow = '\uED41'
	// Soft accent-staccato above
	ArticSoftAccentStaccatoAbove = '\uED42'
	// Soft accent-staccato below
	ArticSoftAccentStaccatoBelow = '\uED43'
	// Soft accent-tenuto above
	ArticSoftAccentTenutoAbove = '\uED44'
	// Soft accent-tenuto below
	ArticSoftAccentTenutoBelow = '\uED45'
	// Soft accent-tenuto-staccato above
	ArticSoftAccentTenutoStaccatoAbove = '\uED46'
	// Soft accent-tenuto-staccato below
	ArticSoftAccentTenutoStaccatoBelow = '\uED47'
	// Staccatissimo above
	ArticStaccatissimoAbove = '\uE4A6'
	// Staccatissimo below
	ArticStaccatissimoBelow = '\uE4A7'
	// Staccatissimo stroke above
	ArticStaccatissimoStrokeAbove = '\uE4AA'
	// Staccatissimo stroke below
	ArticStaccatissimoStrokeBelow = '\uE4AB'
	// Staccatissimo wedge above
	ArticStaccatissimoWedgeAbove = '\uE4A8'
	// Staccatissimo wedge below
	ArticStaccatissimoWedgeBelow = '\uE4A9'
	// Staccato above
	ArticStaccatoAbove = '\uE4A2'
	// Staccato below
	ArticStaccatoBelow = '\uE4A3'
	// Stress above
	ArticStressAbove = '\uE4B6'
	// Stress below
	ArticStressBelow = '\uE4B7'
	// Tenuto above
	ArticTenutoAbove = '\uE4A4'
	// Tenuto-accent above
	ArticTenutoAccentAbove = '\uE4B4'
	// Tenuto-accent below
	ArticTenutoAccentBelow = '\uE4B5'
	// Tenuto below
	ArticTenutoBelow = '\uE4A5'
	// Louré (tenuto-staccato) above
	ArticTenutoStaccatoAbove = '\uE4B2'
	// Louré (tenuto-staccato) below
	ArticTenutoStaccatoBelow = '\uE4B3'
	// Unstress above
	ArticUnstressAbove = '\uE4B8'
	// Unstress below
	ArticUnstressBelow = '\uE4B9'
	// Augmentation dot
	AugmentationDot = '\uE1E7'
	// Dashed barline
	BarlineDashed = '\uE036'
	// Dotted barline
	BarlineDotted = '\uE037'
	// Double barline
	BarlineDouble = '\uE031'
	// Final barline
	BarlineFinal = '\uE032'
	// Heavy barline
	BarlineHeavy = '\uE034'
	// Heavy double barline
	BarlineHeavyHeavy = '\uE035'
	// Reverse final barline
	BarlineReverseFinal = '\uE033'
	// Short barline
	BarlineShort = '\uE038'
	// Single barline
	BarlineSingle = '\uE030'
	// Tick barline
	BarlineTick = '\uE039'
	// Accel./rit. beam 1 (widest)
	BeamAccelRit1 = '\uEAF4'
	// Accel./rit. beam 10
	BeamAccelRit10 = '\uEAFD'
	// Accel./rit. beam 11
	BeamAccelRit11 = '\uEAFE'
	// Accel./rit. beam 12
	BeamAccelRit12 = '\uEAFF'
	// Accel./rit. beam 13
	BeamAccelRit13 = '\uEB00'
	// Accel./rit. beam 14
	BeamAccelRit14 = '\uEB01'
	// Accel./rit. beam 15 (narrowest)
	BeamAccelRit15 = '\uEB02'
	// Accel./rit. beam 2
	BeamAccelRit2 = '\uEAF5'
	// Accel./rit. beam 3
	BeamAccelRit3 = '\uEAF6'
	// Accel./rit. beam 4
	BeamAccelRit4 = '\uEAF7'
	// Accel./rit. beam 5
	BeamAccelRit5 = '\uEAF8'
	// Accel./rit. beam 6
	BeamAccelRit6 = '\uEAF9'
	// Accel./rit. beam 7
	BeamAccelRit7 = '\uEAFA'
	// Accel./rit. beam 8
	BeamAccelRit8 = '\uEAFB'
	// Accel./rit. beam 9
	BeamAccelRit9 = '\uEAFC'
	// Accel./rit. beam terminating line
	BeamAccelRitFinal = '\uEB03'
	// Brace
	Brace = '\uE000'
	// Bracket
	Bracket = '\uE002'
	// Bracket bottom
	BracketBottom = '\uE004'
	// Bracket top
	BracketTop = '\uE003'
	// Bend
	BrassBend = '\uE5E3'
	// Doit, long
	BrassDoitLong = '\uE5D6'
	// Doit, medium
	BrassDoitMedium = '\uE5D5'
	// Doit, short
	BrassDoitShort = '\uE5D4'
	// Lip fall, long
	BrassFallLipLong = '\uE5D9'
	// Lip fall, medium
	BrassFallLipMedium = '\uE5D8'
	// Lip fall, short
	BrassFallLipShort = '\uE5D7'
	// Rough fall, long
	BrassFallRoughLong = '\uE5DF'
	// Rough fall, medium
	BrassFallRoughMedium = '\uE5DE'
	// Rough fall, short
	BrassFallRoughShort = '\uE5DD'
	// Smooth fall, long
	BrassFallSmoothLong = '\uE5DC'
	// Smooth fall, medium
	BrassFallSmoothMedium = '\uE5DB'
	// Smooth fall, short
	BrassFallSmoothShort = '\uE5DA'
	// Flip
	BrassFlip = '\uE5E1'
	// Harmon mute, stem in
	BrassHarmonMuteClosed = '\uE5E8'
	// Harmon mute, stem extended, left
	BrassHarmonMuteStemHalfLeft = '\uE5E9'
	// Harmon mute, stem extended, right
	BrassHarmonMuteStemHalfRight = '\uE5EA'
	// Harmon mute, stem out
	BrassHarmonMuteStemOpen = '\uE5EB'
	// Jazz turn
	BrassJazzTurn = '\uE5E4'
	// Lift, long
	BrassLiftLong = '\uE5D3'
	// Lift, medium
	BrassLiftMedium = '\uE5D2'
	// Lift, short
	BrassLiftShort = '\uE5D1'
	// Smooth lift, long
	BrassLiftSmoothLong = '\uE5EE'
	// Smooth lift, medium
	BrassLiftSmoothMedium = '\uE5ED'
	// Smooth lift, short
	BrassLiftSmoothShort = '\uE5EC'
	// Muted (closed)
	BrassMuteClosed = '\uE5E5'
	// Half-muted (half-closed)
	BrassMuteHalfClosed = '\uE5E6'
	// Open
	BrassMuteOpen = '\uE5E7'
	// Plop
	BrassPlop = '\uE5E0'
	// Scoop
	BrassScoop = '\uE5D0'
	// Smear
	BrassSmear = '\uE5E2'
	// Valve trill
	BrassValveTrill = '\uE5EF'
	// Breath mark (comma)
	BreathMarkComma = '\uE4CE'
	// Breath mark (Salzedo)
	BreathMarkSalzedo = '\uE4D5'
	// Breath mark (tick-like)
	BreathMarkTick = '\uE4CF'
	// Breath mark (upbow-like)
	BreathMarkUpbow = '\uE4D0'
	// Bridge clef
	BridgeClef = '\uE078'
	// Buzz roll
	BuzzRoll = '\uE22A'
	// C clef
	CClef = '\uE05C'
	// C clef ottava bassa
	CClef8vb = '\uE05D'
	// C clef, arrow down
	CClefArrowDown = '\uE05F'
	// C clef, arrow up
	CClefArrowUp = '\uE05E'
	// C clef change
	CClefChange = '\uE07B'
	// Combining C clef
	CClefCombining = '\uE061'
	// Reversed C clef
	CClefReversed = '\uE075'
	// C clef (19th century)
	CClefSquare = '\uE060'
	// Caesura
	Caesura = '\uE4D1'
	// Curved caesura
	CaesuraCurved = '\uE4D4'
	// Short caesura
	CaesuraShort = '\uE4D3'
	// Single stroke caesura
	CaesuraSingleStroke = '\uE4D7'
	// Thick caesura
	CaesuraThick = '\uE4D2'
	// Accentus above
	ChantAccentusAbove = '\uE9D6'
	// Accentus below
	ChantAccentusBelow = '\uE9D7'
	// Punctum auctum, ascending
	ChantAuctumAsc = '\uE994'
	// Punctum auctum, descending
	ChantAuctumDesc = '\uE995'
	// Augmentum (mora)
	ChantAugmentum = '\uE9D9'
	// Caesura
	ChantCaesura = '\uE8F8'
	// Plainchant C clef
	ChantCclef = '\uE906'
	// Circulus above
	ChantCirculusAbove = '\uE9D2'
	// Circulus below
	ChantCirculusBelow = '\uE9D3'
	// Connecting line, ascending 2nd
	ChantConnectingLineAsc2nd = '\uE9BD'
	// Connecting line, ascending 3rd
	ChantConnectingLineAsc3rd = '\uE9BE'
	// Connecting line, ascending 4th
	ChantConnectingLineAsc4th = '\uE9BF'
	// Connecting line, ascending 5th
	ChantConnectingLineAsc5th = '\uE9C0'
	// Connecting line, ascending 6th
	ChantConnectingLineAsc6th = '\uE9C1'
	// Plainchant custos, stem down, high position
	ChantCustosStemDownPosHigh = '\uEA08'
	// Plainchant custos, stem down, highest position
	ChantCustosStemDownPosHighest = '\uEA09'
	// Plainchant custos, stem down, middle position
	ChantCustosStemDownPosMiddle = '\uEA07'
	// Plainchant custos, stem up, low position
	ChantCustosStemUpPosLow = '\uEA05'
	// Plainchant custos, stem up, lowest position
	ChantCustosStemUpPosLowest = '\uEA04'
	// Plainchant custos, stem up, middle position
	ChantCustosStemUpPosMiddle = '\uEA06'
	// Punctum deminutum, lower
	ChantDeminutumLower = '\uE9B3'
	// Punctum deminutum, upper
	ChantDeminutumUpper = '\uE9B2'
	// Divisio finalis
	ChantDivisioFinalis = '\uE8F6'
	// Divisio maior
	ChantDivisioMaior = '\uE8F4'
	// Divisio maxima
	ChantDivisioMaxima = '\uE8F5'
	// Divisio minima
	ChantDivisioMinima = '\uE8F3'
	// Entry line, ascending 2nd
	ChantEntryLineAsc2nd = '\uE9B4'
	// Entry line, ascending 3rd
	ChantEntryLineAsc3rd = '\uE9B5'
	// Entry line, ascending 4th
	ChantEntryLineAsc4th = '\uE9B6'
	// Entry line, ascending 5th
	ChantEntryLineAsc5th = '\uE9B7'
	// Entry line, ascending 6th
	ChantEntryLineAsc6th = '\uE9B8'
	// Episema
	ChantEpisema = '\uE9D8'
	// Plainchant F clef
	ChantFclef = '\uE902'
	// Ictus above
	ChantIctusAbove = '\uE9D0'
	// Ictus below
	ChantIctusBelow = '\uE9D1'
	// Ligated stroke, descending 2nd
	ChantLigaturaDesc2nd = '\uE9B9'
	// Ligated stroke, descending 3rd
	ChantLigaturaDesc3rd = '\uE9BA'
	// Ligated stroke, descending 4th
	ChantLigaturaDesc4th = '\uE9BB'
	// Ligated stroke, descending 5th
	ChantLigaturaDesc5th = '\uE9BC'
	// Oriscus ascending
	ChantOriscusAscending = '\uE99C'
	// Oriscus descending
	ChantOriscusDescending = '\uE99D'
	// Oriscus liquescens
	ChantOriscusLiquescens = '\uE99E'
	// Podatus, lower
	ChantPodatusLower = '\uE9B0'
	// Podatus, upper
	ChantPodatusUpper = '\uE9B1'
	// Punctum
	ChantPunctum = '\uE990'
	// Punctum cavum
	ChantPunctumCavum = '\uE998'
	// Punctum deminutum
	ChantPunctumDeminutum = '\uE9A1'
	// Punctum inclinatum
	ChantPunctumInclinatum = '\uE991'
	// Punctum inclinatum auctum
	ChantPunctumInclinatumAuctum = '\uE992'
	// Punctum inclinatum deminutum
	ChantPunctumInclinatumDeminutum = '\uE993'
	// Punctum linea
	ChantPunctumLinea = '\uE999'
	// Punctum linea cavum
	ChantPunctumLineaCavum = '\uE99A'
	// Punctum virga
	ChantPunctumVirga = '\uE996'
	// Punctum virga, reversed
	ChantPunctumVirgaReversed = '\uE997'
	// Quilisma
	ChantQuilisma = '\uE99B'
	// Semicirculus above
	ChantSemicirculusAbove = '\uE9D4'
	// Semicirculus below
	ChantSemicirculusBelow = '\uE9D5'
	// Plainchant staff
	ChantStaff = '\uE8F0'
	// Plainchant staff (narrow)
	ChantStaffNarrow = '\uE8F2'
	// Plainchant staff (wide)
	ChantStaffWide = '\uE8F1'
	// Strophicus
	ChantStrophicus = '\uE99F'
	// Strophicus auctus
	ChantStrophicusAuctus = '\uE9A0'
	// Strophicus liquescens, 2nd
	ChantStrophicusLiquescens2nd = '\uE9C2'
	// Strophicus liquescens, 3rd
	ChantStrophicusLiquescens3rd = '\uE9C3'
	// Strophicus liquescens, 4th
	ChantStrophicusLiquescens4th = '\uE9C4'
	// Strophicus liquescens, 5th
	ChantStrophicusLiquescens5th = '\uE9C5'
	// Virgula
	ChantVirgula = '\uE8F7'
	// 15 for clefs
	Clef15 = '\uE07E'
	// 8 for clefs
	Clef8 = '\uE07D'
	// Combining clef change
	ClefChangeCombining = '\uE07F'
	// Coda
	Coda = '\uE048'
	// Square coda
	CodaSquare = '\uE049'
	// Beat 2, compound time
	ConductorBeat2Compound = '\uE897'
	// Beat 2, simple time
	ConductorBeat2Simple = '\uE894'
	// Beat 3, compound time
	ConductorBeat3Compound = '\uE898'
	// Beat 3, simple time
	ConductorBeat3Simple = '\uE895'
	// Beat 4, compound time
	ConductorBeat4Compound = '\uE899'
	// Beat 4, simple time
	ConductorBeat4Simple = '\uE896'
	// Left-hand beat or cue
	ConductorLeftBeat = '\uE891'
	// Right-hand beat or cue
	ConductorRightBeat = '\uE892'
	// Strong beat or cue
	ConductorStrongBeat = '\uE890'
	// Unconducted/free passages
	ConductorUnconducted = '\uE89A'
	// Weak beat or cue
	ConductorWeakBeat = '\uE893'
	// Begin beam
	ControlBeginBeam = '\uE8E0'
	// Begin phrase
	ControlBeginPhrase = '\uE8E6'
	// Begin slur
	ControlBeginSlur = '\uE8E4'
	// Begin tie
	ControlBeginTie = '\uE8E2'
	// End beam
	ControlEndBeam = '\uE8E1'
	// End phrase
	ControlEndPhrase = '\uE8E7'
	// End slur
	ControlEndSlur = '\uE8E5'
	// End tie
	ControlEndTie = '\uE8E3'
	// Double flat
	CsymAccidentalDoubleFlat = '\uED64'
	// Double sharp
	CsymAccidentalDoubleSharp = '\uED63'
	// Flat
	CsymAccidentalFlat = '\uED60'
	// Natural
	CsymAccidentalNatural = '\uED61'
	// Sharp
	CsymAccidentalSharp = '\uED62'
	// Triple flat
	CsymAccidentalTripleFlat = '\uED66'
	// Triple sharp
	CsymAccidentalTripleSharp = '\uED65'
	// Slash for altered bass note
	CsymAlteredBassSlash = '\uE87B'
	// Augmented
	CsymAugmented = '\uE872'
	// Double-height left bracket
	CsymBracketLeftTall = '\uE877'
	// Double-height right bracket
	CsymBracketRightTall = '\uE878'
	// Slash for chord symbols arranged diagonally
	CsymDiagonalArrangementSlash = '\uE87C'
	// Diminished
	CsymDiminished = '\uE870'
	// Half-diminished
	CsymHalfDiminished = '\uE871'
	// Major seventh
	CsymMajorSeventh = '\uE873'
	// Minor
	CsymMinor = '\uE874'
	// Double-height left parenthesis
	CsymParensLeftTall = '\uE875'
	// Triple-height left parenthesis
	CsymParensLeftVeryTall = '\uE879'
	// Double-height right parenthesis
	CsymParensRightTall = '\uE876'
	// Triple-height right parenthesis
	CsymParensRightVeryTall = '\uE87A'
	// Curlew (Britten)
	CurlewSign = '\uE4D6'
	// Da capo
	DaCapo = '\uE046'
	// Dal segno
	DalSegno = '\uE045'
	// Daseian excellentes 1
	DaseianExcellentes1 = '\uEA3C'
	// Daseian excellentes 2
	DaseianExcellentes2 = '\uEA3D'
	// Daseian excellentes 3
	DaseianExcellentes3 = '\uEA3E'
	// Daseian excellentes 4
	DaseianExcellentes4 = '\uEA3F'
	// Daseian finales 1
	DaseianFinales1 = '\uEA34'
	// Daseian finales 2
	DaseianFinales2 = '\uEA35'
	// Daseian finales 3
	DaseianFinales3 = '\uEA36'
	// Daseian finales 4
	DaseianFinales4 = '\uEA37'
	// Daseian graves 1
	DaseianGraves1 = '\uEA30'
	// Daseian graves 2
	DaseianGraves2 = '\uEA31'
	// Daseian graves 3
	DaseianGraves3 = '\uEA32'
	// Daseian graves 4
	DaseianGraves4 = '\uEA33'
	// Daseian residua 1
	DaseianResidua1 = '\uEA40'
	// Daseian residua 2
	DaseianResidua2 = '\uEA41'
	// Daseian superiores 1
	DaseianSuperiores1 = '\uEA38'
	// Daseian superiores 2
	DaseianSuperiores2 = '\uEA39'
	// Daseian superiores 3
	DaseianSuperiores3 = '\uEA3A'
	// Daseian superiores 4
	DaseianSuperiores4 = '\uEA3B'
	// Double lateral roll (Stevens)
	DoubleLateralRollStevens = '\uE234'
	// Double-tongue above
	DoubleTongueAbove = '\uE5F0'
	// Double-tongue below
	DoubleTongueBelow = '\uE5F1'
	// Colon separator for combined dynamics
	DynamicCombinedSeparatorColon = '\uE546'
	// Hyphen separator for combined dynamics
	DynamicCombinedSeparatorHyphen = '\uE547'
	// Slash separator for combined dynamics
	DynamicCombinedSeparatorSlash = '\uE549'
	// Space separator for combined dynamics
	DynamicCombinedSeparatorSpace = '\uE548'
	// Crescendo
	DynamicCrescendoHairpin = '\uE53E'
	// Diminuendo
	DynamicDiminuendoHairpin = '\uE53F'
	// ff
	DynamicFF = '\uE52F'
	// fff
	DynamicFFF = '\uE530'
	// ffff
	DynamicFFFF = '\uE531'
	// fffff
	DynamicFFFFF = '\uE532'
	// ffffff
	DynamicFFFFFF = '\uE533'
	// Forte
	DynamicForte = '\uE522'
	// Forte-piano
	DynamicFortePiano = '\uE534'
	// Forzando
	DynamicForzando = '\uE535'
	// Left bracket (for hairpins)
	DynamicHairpinBracketLeft = '\uE544'
	// Right bracket (for hairpins)
	DynamicHairpinBracketRight = '\uE545'
	// Left parenthesis (for hairpins)
	DynamicHairpinParenthesisLeft = '\uE542'
	// Right parenthesis (for hairpins)
	DynamicHairpinParenthesisRight = '\uE543'
	// mf
	DynamicMF = '\uE52D'
	// mp
	DynamicMP = '\uE52C'
	// Messa di voce
	DynamicMessaDiVoce = '\uE540'
	// Mezzo
	DynamicMezzo = '\uE521'
	// Niente
	DynamicNiente = '\uE526'
	// Niente (for hairpins)
	DynamicNienteForHairpin = '\uE541'
	// pf
	DynamicPF = '\uE52E'
	// pp
	DynamicPP = '\uE52B'
	// ppp
	DynamicPPP = '\uE52A'
	// pppp
	DynamicPPPP = '\uE529'
	// ppppp
	DynamicPPPPP = '\uE528'
	// pppppp
	DynamicPPPPPP = '\uE527'
	// Piano
	DynamicPiano = '\uE520'
	// Rinforzando
	DynamicRinforzando = '\uE523'
	// Rinforzando 1
	DynamicRinforzando1 = '\uE53C'
	// Rinforzando 2
	DynamicRinforzando2 = '\uE53D'
	// Sforzando
	DynamicSforzando = '\uE524'
	// Sforzando 1
	DynamicSforzando1 = '\uE536'
	// Sforzando-pianissimo
	DynamicSforzandoPianissimo = '\uE538'
	// Sforzando-piano
	DynamicSforzandoPiano = '\uE537'
	// Sforzato
	DynamicSforzato = '\uE539'
	// Sforzatissimo
	DynamicSforzatoFF = '\uE53B'
	// Sforzato-piano
	DynamicSforzatoPiano = '\uE53A'
	// Z
	DynamicZ = '\uE525'
	// Eight channels (7.1 surround)
	ElecAudioChannelsEight = '\uEB46'
	// Five channels
	ElecAudioChannelsFive = '\uEB43'
	// Four channels
	ElecAudioChannelsFour = '\uEB42'
	// One channel (mono)
	ElecAudioChannelsOne = '\uEB3E'
	// Seven channels
	ElecAudioChannelsSeven = '\uEB45'
	// Six channels (5.1 surround)
	ElecAudioChannelsSix = '\uEB44'
	// Three channels (frontal)
	ElecAudioChannelsThreeFrontal = '\uEB40'
	// Three channels (surround)
	ElecAudioChannelsThreeSurround = '\uEB41'
	// Two channels (stereo)
	ElecAudioChannelsTwo = '\uEB3F'
	// Audio in
	ElecAudioIn = '\uEB49'
	// Mono audio setup
	ElecAudioMono = '\uEB3C'
	// Audio out
	ElecAudioOut = '\uEB4A'
	// Stereo audio setup
	ElecAudioStereo = '\uEB3D'
	// Camera
	ElecCamera = '\uEB1B'
	// Data in
	ElecDataIn = '\uEB4D'
	// Data out
	ElecDataOut = '\uEB4E'
	// Disc
	ElecDisc = '\uEB13'
	// Download
	ElecDownload = '\uEB4F'
	// Eject
	ElecEject = '\uEB2B'
	// Fast-forward
	ElecFastForward = '\uEB1F'
	// Headphones
	ElecHeadphones = '\uEB11'
	// Headset
	ElecHeadset = '\uEB12'
	// Line in
	ElecLineIn = '\uEB47'
	// Line out
	ElecLineOut = '\uEB48'
	// Loop
	ElecLoop = '\uEB23'
	// Loudspeaker
	ElecLoudspeaker = '\uEB1A'
	// MIDI controller 0%
	ElecMIDIController0 = '\uEB36'
	// MIDI controller 100%
	ElecMIDIController100 = '\uEB3B'
	// MIDI controller 20%
	ElecMIDIController20 = '\uEB37'
	// MIDI controller 40%
	ElecMIDIController40 = '\uEB38'
	// MIDI controller 60%
	ElecMIDIController60 = '\uEB39'
	// MIDI controller 80%
	ElecMIDIController80 = '\uEB3A'
	// MIDI in
	ElecMIDIIn = '\uEB34'
	// MIDI out
	ElecMIDIOut = '\uEB35'
	// Microphone
	ElecMicrophone = '\uEB10'
	// Mute microphone
	ElecMicrophoneMute = '\uEB28'
	// Unmute microphone
	ElecMicrophoneUnmute = '\uEB29'
	// Mixing console
	ElecMixingConsole = '\uEB15'
	// Monitor
	ElecMonitor = '\uEB18'
	// Mute
	ElecMute = '\uEB26'
	// Pause
	ElecPause = '\uEB1E'
	// Play
	ElecPlay = '\uEB1C'
	// Power on/off
	ElecPowerOnOff = '\uEB2A'
	// Projector
	ElecProjector = '\uEB19'
	// Replay
	ElecReplay = '\uEB24'
	// Rewind
	ElecRewind = '\uEB20'
	// Shuffle
	ElecShuffle = '\uEB25'
	// Skip backwards
	ElecSkipBackwards = '\uEB22'
	// Skip forwards
	ElecSkipForwards = '\uEB21'
	// Stop
	ElecStop = '\uEB1D'
	// Tape
	ElecTape = '\uEB14'
	// USB connection
	ElecUSB = '\uEB16'
	// Unmute
	ElecUnmute = '\uEB27'
	// Upload
	ElecUpload = '\uEB50'
	// Video camera
	ElecVideoCamera = '\uEB17'
	// Video in
	ElecVideoIn = '\uEB4B'
	// Video out
	ElecVideoOut = '\uEB4C'
	// Combining volume fader
	ElecVolumeFader = '\uEB2C'
	// Combining volume fader thumb
	ElecVolumeFaderThumb = '\uEB2D'
	// Volume level 0%
	ElecVolumeLevel0 = '\uEB2E'
	// Volume level 100%
	ElecVolumeLevel100 = '\uEB33'
	// Volume level 20%
	ElecVolumeLevel20 = '\uEB2F'
	// Volume level 40%
	ElecVolumeLevel40 = '\uEB30'
	// Volume level 60%
	ElecVolumeLevel60 = '\uEB31'
	// Volume level 80%
	ElecVolumeLevel80 = '\uEB32'
	// F clef
	FClef = '\uE062'
	// F clef quindicesima alta
	FClef15ma = '\uE066'
	// F clef quindicesima bassa
	FClef15mb = '\uE063'
	// F clef ottava alta
	FClef8va = '\uE065'
	// F clef ottava bassa
	FClef8vb = '\uE064'
	// F clef, arrow down
	FClefArrowDown = '\uE068'
	// F clef, arrow up
	FClefArrowUp = '\uE067'
	// F clef change
	FClefChange = '\uE07C'
	// Reversed F clef
	FClefReversed = '\uE076'
	// Turned F clef
	FClefTurned = '\uE077'
	// Fermata above
	FermataAbove = '\uE4C0'
	// Fermata below
	FermataBelow = '\uE4C1'
	// Long fermata above
	FermataLongAbove = '\uE4C6'
	// Long fermata below
	FermataLongBelow = '\uE4C7'
	// Long fermata (Henze) above
	FermataLongHenzeAbove = '\uE4CA'
	// Long fermata (Henze) below
	FermataLongHenzeBelow = '\uE4CB'
	// Short fermata above
	FermataShortAbove = '\uE4C4'
	// Short fermata below
	FermataShortBelow = '\uE4C5'
	// Short fermata (Henze) above
	FermataShortHenzeAbove = '\uE4CC'
	// Short fermata (Henze) below
	FermataShortHenzeBelow = '\uE4CD'
	// Very long fermata above
	FermataVeryLongAbove = '\uE4C8'
	// Very long fermata below
	FermataVeryLongBelow = '\uE4C9'
	// Very short fermata above
	FermataVeryShortAbove = '\uE4C2'
	// Very short fermata below
	FermataVeryShortBelow = '\uE4C3'
	// Figured bass 0
	Figbass0 = '\uEA50'
	// Figured bass 1
	Figbass1 = '\uEA51'
	// Figured bass 2
	Figbass2 = '\uEA52'
	// Figured bass 2 raised by half-step
	Figbass2Raised = '\uEA53'
	// Figured bass 3
	Figbass3 = '\uEA54'
	// Figured bass 4
	Figbass4 = '\uEA55'
	// Figured bass 4 raised by half-step
	Figbass4Raised = '\uEA56'
	// Figured bass 5
	Figbass5 = '\uEA57'
	// Figured bass 5 raised by half-step
	Figbass5Raised1 = '\uEA58'
	// Figured bass 5 raised by half-step 2
	Figbass5Raised2 = '\uEA59'
	// Figured bass diminished 5
	Figbass5Raised3 = '\uEA5A'
	// Figured bass 6
	Figbass6 = '\uEA5B'
	// Figured bass 6 raised by half-step
	Figbass6Raised = '\uEA5C'
	// Figured bass 6 raised by half-step 2
	Figbass6Raised2 = '\uEA6F'
	// Figured bass 7
	Figbass7 = '\uEA5D'
	// Figured bass 7 diminished
	Figbass7Diminished = '\uECC0'
	// Figured bass 7 raised by half-step
	Figbass7Raised1 = '\uEA5E'
	// Figured bass 7 lowered by a half-step
	Figbass7Raised2 = '\uEA5F'
	// Figured bass 8
	Figbass8 = '\uEA60'
	// Figured bass 9
	Figbass9 = '\uEA61'
	// Figured bass 9 raised by half-step
	Figbass9Raised = '\uEA62'
	// Figured bass [
	FigbassBracketLeft = '\uEA68'
	// Figured bass ]
	FigbassBracketRight = '\uEA69'
	// Combining lower
	FigbassCombiningLowering = '\uEA6E'
	// Combining raise
	FigbassCombiningRaising = '\uEA6D'
	// Figured bass double flat
	FigbassDoubleFlat = '\uEA63'
	// Figured bass double sharp
	FigbassDoubleSharp = '\uEA67'
	// Figured bass flat
	FigbassFlat = '\uEA64'
	// Figured bass natural
	FigbassNatural = '\uEA65'
	// Figured bass (
	FigbassParensLeft = '\uEA6A'
	// Figured bass )
	FigbassParensRight = '\uEA6B'
	// Figured bass +
	FigbassPlus = '\uEA6C'
	// Figured bass sharp
	FigbassSharp = '\uEA66'
	// Figured bass triple flat
	FigbassTripleFlat = '\uECC1'
	// Figured bass triple sharp
	FigbassTripleSharp = '\uECC2'
	// Fingering 0 (open string)
	Fingering0 = '\uED10'
	// Fingering 0 italic (open string)
	Fingering0Italic = '\uED80'
	// Fingering 1 (thumb)
	Fingering1 = '\uED11'
	// Fingering 1 italic (thumb)
	Fingering1Italic = '\uED81'
	// Fingering 2 (index finger)
	Fingering2 = '\uED12'
	// Fingering 2 italic (index finger)
	Fingering2Italic = '\uED82'
	// Fingering 3 (middle finger)
	Fingering3 = '\uED13'
	// Fingering 3 italic (middle finger)
	Fingering3Italic = '\uED83'
	// Fingering 4 (ring finger)
	Fingering4 = '\uED14'
	// Fingering 4 italic (ring finger)
	Fingering4Italic = '\uED84'
	// Fingering 5 (little finger)
	Fingering5 = '\uED15'
	// Fingering 5 italic (little finger)
	Fingering5Italic = '\uED85'
	// Fingering 6
	Fingering6 = '\uED24'
	// Fingering 6 italic
	Fingering6Italic = '\uED86'
	// Fingering 7
	Fingering7 = '\uED25'
	// Fingering 7 italic
	Fingering7Italic = '\uED87'
	// Fingering 8
	Fingering8 = '\uED26'
	// Fingering 8 italic
	Fingering8Italic = '\uED88'
	// Fingering 9
	Fingering9 = '\uED27'
	// Fingering 9 italic
	Fingering9Italic = '\uED89'
	// Fingering a (anular; right-hand ring finger for guitar)
	FingeringALower = '\uED1B'
	// Fingering c (right-hand little finger for guitar)
	FingeringCLower = '\uED1C'
	// Fingering e (right-hand little finger for guitar)
	FingeringELower = '\uED1E'
	// Fingering i (indicio; right-hand index finger for guitar)
	FingeringILower = '\uED19'
	// Fingering left bracket
	FingeringLeftBracket = '\uED2A'
	// Fingering left bracket italic
	FingeringLeftBracketItalic = '\uED8C'
	// Fingering left parenthesis
	FingeringLeftParenthesis = '\uED28'
	// Fingering left parenthesis italic
	FingeringLeftParenthesisItalic = '\uED8A'
	// Fingering m (medio; right-hand middle finger for guitar)
	FingeringMLower = '\uED1A'
	// Multiple notes played by thumb or single finger
	FingeringMultipleNotes = '\uED23'
	// Fingering o (right-hand little finger for guitar)
	FingeringOLower = '\uED1F'
	// Fingering p (pulgar; right-hand thumb for guitar)
	FingeringPLower = '\uED17'
	// Fingering q (right-hand little finger for guitar)
	FingeringQLower = '\uED8E'
	// Fingering right bracket
	FingeringRightBracket = '\uED2B'
	// Fingering right bracket italic
	FingeringRightBracketItalic = '\uED8D'
	// Fingering right parenthesis
	FingeringRightParenthesis = '\uED29'
	// Fingering right parenthesis italic
	FingeringRightParenthesisItalic = '\uED8B'
	// Fingering s (right-hand little finger for guitar)
	FingeringSLower = '\uED8F'
	// Fingering middle dot separator
	FingeringSeparatorMiddleDot = '\uED2C'
	// Fingering white middle dot separator
	FingeringSeparatorMiddleDotWhite = '\uED2D'
	// Fingering forward slash separator
	FingeringSeparatorSlash = '\uED2E'
	// Finger substitution above
	FingeringSubstitutionAbove = '\uED20'
	// Finger substitution below
	FingeringSubstitutionBelow = '\uED21'
	// Finger substitution dash
	FingeringSubstitutionDash = '\uED22'
	// Fingering t (right-hand thumb for guitar)
	FingeringTLower = '\uED18'
	// Fingering T (left-hand thumb for guitar)
	FingeringTUpper = '\uED16'
	// Fingering x (right-hand little finger for guitar)
	FingeringXLower = '\uED1D'
	// Combining flag 8 (1024th) below
	Flag1024thDown = '\uE24F'
	// Combining flag 8 (1024th) above
	Flag1024thUp = '\uE24E'
	// Combining flag 5 (128th) below
	Flag128thDown = '\uE249'
	// Combining flag 5 (128th) above
	Flag128thUp = '\uE248'
	// Combining flag 2 (16th) below
	Flag16thDown = '\uE243'
	// Combining flag 2 (16th) above
	Flag16thUp = '\uE242'
	// Combining flag 6 (256th) below
	Flag256thDown = '\uE24B'
	// Combining flag 6 (256th) above
	Flag256thUp = '\uE24A'
	// Combining flag 3 (32nd) below
	Flag32ndDown = '\uE245'
	// Combining flag 3 (32nd) above
	Flag32ndUp = '\uE244'
	// Combining flag 7 (512th) below
	Flag512thDown = '\uE24D'
	// Combining flag 7 (512th) above
	Flag512thUp = '\uE24C'
	// Combining flag 4 (64th) below
	Flag64thDown = '\uE247'
	// Combining flag 4 (64th) above
	Flag64thUp = '\uE246'
	// Combining flag 1 (8th) below
	Flag8thDown = '\uE241'
	// Combining flag 1 (8th) above
	Flag8thUp = '\uE240'
	// Internal combining flag below
	FlagInternalDown = '\uE251'
	// Internal combining flag above
	FlagInternalUp = '\uE250'
	// 3-string fretboard
	Fretboard3String = '\uE850'
	// 3-string fretboard at nut
	Fretboard3StringNut = '\uE851'
	// 4-string fretboard
	Fretboard4String = '\uE852'
	// 4-string fretboard at nut
	Fretboard4StringNut = '\uE853'
	// 5-string fretboard
	Fretboard5String = '\uE854'
	// 5-string fretboard at nut
	Fretboard5StringNut = '\uE855'
	// 6-string fretboard
	Fretboard6String = '\uE856'
	// 6-string fretboard at nut
	Fretboard6StringNut = '\uE857'
	// Fingered fret (filled circle)
	FretboardFilledCircle = '\uE858'
	// Open string (O)
	FretboardO = '\uE85A'
	// String not played (X)
	FretboardX = '\uE859'
	// Function theory angle bracket left
	FunctionAngleLeft = '\uEA93'
	// Function theory angle bracket right
	FunctionAngleRight = '\uEA94'
	// Function theory bracket left
	FunctionBracketLeft = '\uEA8F'
	// Function theory bracket right
	FunctionBracketRight = '\uEA90'
	// Function theory dominant of dominant
	FunctionDD = '\uEA81'
	// Function theory minor dominant
	FunctionDLower = '\uEA80'
	// Function theory major dominant
	FunctionDUpper = '\uEA7F'
	// Function theory 8
	FunctionEight = '\uEA78'
	// Function theory F
	FunctionFUpper = '\uEA99'
	// Function theory 5
	FunctionFive = '\uEA75'
	// Function theory 4
	FunctionFour = '\uEA74'
	// Function theory g
	FunctionGLower = '\uEA84'
	// Function theory G
	FunctionGUpper = '\uEA83'
	// Function theory greater than
	FunctionGreaterThan = '\uEA7C'
	// Function theory i
	FunctionILower = '\uEA9B'
	// Function theory I
	FunctionIUpper = '\uEA9A'
	// Function theory k
	FunctionKLower = '\uEA9D'
	// Function theory K
	FunctionKUpper = '\uEA9C'
	// Function theory l
	FunctionLLower = '\uEA9F'
	// Function theory L
	FunctionLUpper = '\uEA9E'
	// Function theory less than
	FunctionLessThan = '\uEA7A'
	// Function theory m
	FunctionMLower = '\uED01'
	// Function theory M
	FunctionMUpper = '\uED00'
	// Function theory minus
	FunctionMinus = '\uEA7B'
	// Function theory n
	FunctionNLower = '\uEA86'
	// Function theory N
	FunctionNUpper = '\uEA85'
	// Function theory superscript N
	FunctionNUpperSuperscript = '\uED02'
	// Function theory 9
	FunctionNine = '\uEA79'
	// Function theory 1
	FunctionOne = '\uEA71'
	// Function theory p
	FunctionPLower = '\uEA88'
	// Function theory P
	FunctionPUpper = '\uEA87'
	// Function theory parenthesis left
	FunctionParensLeft = '\uEA91'
	// Function theory parenthesis right
	FunctionParensRight = '\uEA92'
	// Function theory prefix plus
	FunctionPlus = '\uEA98'
	// Function theory r
	FunctionRLower = '\uED03'
	// Function theory repetition 1
	FunctionRepetition1 = '\uEA95'
	// Function theory repetition 2
	FunctionRepetition2 = '\uEA96'
	// Function theory prefix ring
	FunctionRing = '\uEA97'
	// Function theory minor subdominant
	FunctionSLower = '\uEA8A'
	// Function theory minor subdominant of subdominant
	FunctionSSLower = '\uEA7E'
	// Function theory major subdominant of subdominant
	FunctionSSUpper = '\uEA7D'
	// Function theory major subdominant
	FunctionSUpper = '\uEA89'
	// Function theory 7
	FunctionSeven = '\uEA77'
	// Function theory 6
	FunctionSix = '\uEA76'
	// Function theory double dominant seventh
	FunctionSlashedDD = '\uEA82'
	// Function theory minor tonic
	FunctionTLower = '\uEA8C'
	// Function theory tonic
	FunctionTUpper = '\uEA8B'
	// Function theory 3
	FunctionThree = '\uEA73'
	// Function theory 2
	FunctionTwo = '\uEA72'
	// Function theory v
	FunctionVLower = '\uEA8E'
	// Function theory V
	FunctionVUpper = '\uEA8D'
	// Function theory 0
	FunctionZero = '\uEA70'
	// G clef
	GClef = '\uE050'
	// G clef quindicesima alta
	GClef15ma = '\uE054'
	// G clef quindicesima bassa
	GClef15mb = '\uE051'
	// G clef ottava alta
	GClef8va = '\uE053'
	// G clef ottava bassa
	GClef8vb = '\uE052'
	// G clef ottava bassa with C clef
	GClef8vbCClef = '\uE056'
	// G clef ottava bassa (old style)
	GClef8vbOld = '\uE055'
	// G clef, optionally ottava bassa
	GClef8vbParens = '\uE057'
	// G clef, arrow down
	GClefArrowDown = '\uE05B'
	// G clef, arrow up
	GClefArrowUp = '\uE05A'
	// G clef change
	GClefChange = '\uE07A'
	// Combining G clef, number above
	GClefLigatedNumberAbove = '\uE059'
	// Combining G clef, number below
	GClefLigatedNumberBelow = '\uE058'
	// Reversed G clef
	GClefReversed = '\uE073'
	// Turned G clef
	GClefTurned = '\uE074'
	// Glissando down
	GlissandoDown = '\uE586'
	// Glissando up
	GlissandoUp = '\uE585'
	// Slashed grace note stem down
	GraceNoteAcciaccaturaStemDown = '\uE561'
	// Slashed grace note stem up
	GraceNoteAcciaccaturaStemUp = '\uE560'
	// Grace note stem down
	GraceNoteAppoggiaturaStemDown = '\uE563'
	// Grace note stem up
	GraceNoteAppoggiaturaStemUp = '\uE562'
	// Slash for stem down grace note
	GraceNoteSlashStemDown = '\uE565'
	// Slash for stem up grace note
	GraceNoteSlashStemUp = '\uE564'
	// Full barré
	GuitarBarreFull = '\uE848'
	// Half barré
	GuitarBarreHalf = '\uE849'
	// Closed wah/volume pedal
	GuitarClosePedal = '\uE83F'
	// Fade in
	GuitarFadeIn = '\uE843'
	// Fade out
	GuitarFadeOut = '\uE844'
	// Golpe (tapping the pick guard)
	GuitarGolpe = '\uE842'
	// Half-open wah/volume pedal
	GuitarHalfOpenPedal = '\uE83E'
	// Left-hand tapping
	GuitarLeftHandTapping = '\uE840'
	// Open wah/volume pedal
	GuitarOpenPedal = '\uE83D'
	// Right-hand tapping
	GuitarRightHandTapping = '\uE841'
	// Guitar shake
	GuitarShake = '\uE832'
	// String number 0
	GuitarString0 = '\uE833'
	// String number 1
	GuitarString1 = '\uE834'
	// String number 10
	GuitarString10 = '\uE84A'
	// String number 11
	GuitarString11 = '\uE84B'
	// String number 12
	GuitarString12 = '\uE84C'
	// String number 13
	GuitarString13 = '\uE84D'
	// String number 2
	GuitarString2 = '\uE835'
	// String number 3
	GuitarString3 = '\uE836'
	// String number 4
	GuitarString4 = '\uE837'
	// String number 5
	GuitarString5 = '\uE838'
	// String number 6
	GuitarString6 = '\uE839'
	// String number 7
	GuitarString7 = '\uE83A'
	// String number 8
	GuitarString8 = '\uE83B'
	// String number 9
	GuitarString9 = '\uE83C'
	// Strum direction down
	GuitarStrumDown = '\uE847'
	// Strum direction up
	GuitarStrumUp = '\uE846'
	// Guitar vibrato bar dip
	GuitarVibratoBarDip = '\uE831'
	// Guitar vibrato bar scoop
	GuitarVibratoBarScoop = '\uE830'
	// Vibrato wiggle segment
	GuitarVibratoStroke = '\uEAB2'
	// Volume swell
	GuitarVolumeSwell = '\uE845'
	// Wide vibrato wiggle segment
	GuitarWideVibratoStroke = '\uEAB3'
	// Belltree
	HandbellsBelltree = '\uE81F'
	// Damp 3
	HandbellsDamp3 = '\uE81E'
	// Echo
	HandbellsEcho1 = '\uE81B'
	// Echo 2
	HandbellsEcho2 = '\uE81C'
	// Gyro
	HandbellsGyro = '\uE81D'
	// Hand martellato
	HandbellsHandMartellato = '\uE812'
	// Mallet, bell on table
	HandbellsMalletBellOnTable = '\uE815'
	// Mallet, bell suspended
	HandbellsMalletBellSuspended = '\uE814'
	// Mallet lift
	HandbellsMalletLft = '\uE816'
	// Martellato
	HandbellsMartellato = '\uE810'
	// Martellato lift
	HandbellsMartellatoLift = '\uE811'
	// Muted martellato
	HandbellsMutedMartellato = '\uE813'
	// Pluck lift
	HandbellsPluckLift = '\uE817'
	// Swing
	HandbellsSwing = '\uE81A'
	// Swing down
	HandbellsSwingDown = '\uE819'
	// Swing up
	HandbellsSwingUp = '\uE818'
	// Table pair of handbells
	HandbellsTablePairBells = '\uE821'
	// Table single handbell
	HandbellsTableSingleBell = '\uE820'
	// Metal rod pictogram
	HarpMetalRod = '\uE68F'
	// Harp pedal centered (natural)
	HarpPedalCentered = '\uE681'
	// Harp pedal divider
	HarpPedalDivider = '\uE683'
	// Harp pedal lowered (sharp)
	HarpPedalLowered = '\uE682'
	// Harp pedal raised (flat)
	HarpPedalRaised = '\uE680'
	// Ascending aeolian chords (Salzedo)
	HarpSalzedoAeolianAscending = '\uE695'
	// Descending aeolian chords (Salzedo)
	HarpSalzedoAeolianDescending = '\uE696'
	// Damp above (Salzedo)
	HarpSalzedoDampAbove = '\uE69A'
	// Damp below (Salzedo)
	HarpSalzedoDampBelow = '\uE699'
	// Damp with both hands (Salzedo)
	HarpSalzedoDampBothHands = '\uE698'
	// Damp only low strings (Salzedo)
	HarpSalzedoDampLowStrings = '\uE697'
	// Fluidic sounds, left hand (Salzedo)
	HarpSalzedoFluidicSoundsLeft = '\uE68D'
	// Fluidic sounds, right hand (Salzedo)
	HarpSalzedoFluidicSoundsRight = '\uE68E'
	// Isolated sounds (Salzedo)
	HarpSalzedoIsolatedSounds = '\uE69C'
	// Metallic sounds (Salzedo)
	HarpSalzedoMetallicSounds = '\uE688'
	// Metallic sounds, one string (Salzedo)
	HarpSalzedoMetallicSoundsOneString = '\uE69B'
	// Muffle totally (Salzedo)
	HarpSalzedoMuffleTotally = '\uE68C'
	// Oboic flux (Salzedo)
	HarpSalzedoOboicFlux = '\uE685'
	// Play at upper end of strings (Salzedo)
	HarpSalzedoPlayUpperEnd = '\uE68A'
	// Slide with suppleness (Salzedo)
	HarpSalzedoSlideWithSuppleness = '\uE684'
	// Snare drum effect (Salzedo)
	HarpSalzedoSnareDrum = '\uE69D'
	// Tam-tam sounds (Salzedo)
	HarpSalzedoTamTamSounds = '\uE689'
	// Thunder effect (Salzedo)
	HarpSalzedoThunderEffect = '\uE686'
	// Timpanic sounds (Salzedo)
	HarpSalzedoTimpanicSounds = '\uE68B'
	// Whistling sounds (Salzedo)
	HarpSalzedoWhistlingSounds = '\uE687'
	// Combining string noise for stem
	HarpStringNoiseStem = '\uE694'
	// Tuning key pictogram
	HarpTuningKey = '\uE690'
	// Retune strings for glissando
	HarpTuningKeyGlissando = '\uE693'
	// Use handle of tuning key pictogram
	HarpTuningKeyHandle = '\uE691'
	// Use shank of tuning key pictogram
	HarpTuningKeyShank = '\uE692'
	// Indian drum clef
	IndianDrumClef = '\uED70'
	// Back-chug
	KahnBackChug = '\uEDE2'
	// Back-flap
	KahnBackFlap = '\uEDD8'
	// Back-riff
	KahnBackRiff = '\uEDE1'
	// Back-rip
	KahnBackRip = '\uEDDA'
	// Ball-change
	KahnBallChange = '\uEDC6'
	// Ball-dig
	KahnBallDig = '\uEDCD'
	// Brush-backward
	KahnBrushBackward = '\uEDA7'
	// Brush-forward
	KahnBrushForward = '\uEDA6'
	// Chug
	KahnChug = '\uEDDD'
	// Clap
	KahnClap = '\uEDB8'
	// Double-snap
	KahnDoubleSnap = '\uEDBA'
	// Double-wing
	KahnDoubleWing = '\uEDEB'
	// Draw-step
	KahnDrawStep = '\uEDB2'
	// Draw-tap
	KahnDrawTap = '\uEDB3'
	// Flam
	KahnFlam = '\uEDCF'
	// Flap
	KahnFlap = '\uEDD5'
	// Flap-step
	KahnFlapStep = '\uEDD7'
	// Flat
	KahnFlat = '\uEDA9'
	// Flea-hop
	KahnFleaHop = '\uEDB0'
	// Flea-tap
	KahnFleaTap = '\uEDB1'
	// Grace-tap
	KahnGraceTap = '\uEDA8'
	// Grace-tap-change
	KahnGraceTapChange = '\uEDD1'
	// Grace-tap-hop
	KahnGraceTapHop = '\uEDD0'
	// Grace-tap-stamp
	KahnGraceTapStamp = '\uEDD3'
	// Heel
	KahnHeel = '\uEDAA'
	// Heel-change
	KahnHeelChange = '\uEDC9'
	// Heel-click
	KahnHeelClick = '\uEDBB'
	// Heel-drop
	KahnHeelDrop = '\uEDB6'
	// Heel-step
	KahnHeelStep = '\uEDC4'
	// Heel-tap
	KahnHeelTap = '\uEDCB'
	// Hop
	KahnHop = '\uEDA2'
	// Jump-apart
	KahnJumpApart = '\uEDA5'
	// Jump-together
	KahnJumpTogether = '\uEDA4'
	// Knee-inward
	KahnKneeInward = '\uEDAD'
	// Knee-outward
	KahnKneeOutward = '\uEDAC'
	// Leap
	KahnLeap = '\uEDA3'
	// Leap-flat-foot
	KahnLeapFlatFoot = '\uEDD2'
	// Leap-heel-click
	KahnLeapHeelClick = '\uEDD4'
	// Left-catch
	KahnLeftCatch = '\uEDBF'
	// Left-cross
	KahnLeftCross = '\uEDBD'
	// Left-foot
	KahnLeftFoot = '\uEDEE'
	// Left-toe-strike
	KahnLeftToeStrike = '\uEDC1'
	// Left-turn
	KahnLeftTurn = '\uEDF0'
	// Over-the-top
	KahnOverTheTop = '\uEDEC'
	// Over-the-top-tap
	KahnOverTheTopTap = '\uEDED'
	// Pull
	KahnPull = '\uEDE3'
	// Push
	KahnPush = '\uEDDE'
	// Riff
	KahnRiff = '\uEDE0'
	// Riffle
	KahnRiffle = '\uEDE7'
	// Right-catch
	KahnRightCatch = '\uEDC0'
	// Right-cross
	KahnRightCross = '\uEDBE'
	// Right-foot
	KahnRightFoot = '\uEDEF'
	// Right-toe-strike
	KahnRightToeStrike = '\uEDC2'
	// Right-turn
	KahnRightTurn = '\uEDF1'
	// Rip
	KahnRip = '\uEDD6'
	// Ripple
	KahnRipple = '\uEDE8'
	// Scrape
	KahnScrape = '\uEDAE'
	// Scuff
	KahnScuff = '\uEDDC'
	// Scuffle
	KahnScuffle = '\uEDE6'
	// Shuffle
	KahnShuffle = '\uEDE5'
	// Slam
	KahnSlam = '\uEDCE'
	// Slap
	KahnSlap = '\uEDD9'
	// Slide-step
	KahnSlideStep = '\uEDB4'
	// Slide-tap
	KahnSlideTap = '\uEDB5'
	// Snap
	KahnSnap = '\uEDB9'
	// Stamp
	KahnStamp = '\uEDC3'
	// Stamp-stamp
	KahnStampStamp = '\uEDC8'
	// Step
	KahnStep = '\uEDA0'
	// Step-stamp
	KahnStepStamp = '\uEDC7'
	// Stomp
	KahnStomp = '\uEDCA'
	// Stomp-brush
	KahnStompBrush = '\uEDDB'
	// Tap
	KahnTap = '\uEDA1'
	// Toe
	KahnToe = '\uEDAB'
	// Toe-click
	KahnToeClick = '\uEDBC'
	// Toe-drop
	KahnToeDrop = '\uEDB7'
	// Toe-step
	KahnToeStep = '\uEDC5'
	// Toe-tap
	KahnToeTap = '\uEDCC'
	// Trench
	KahnTrench = '\uEDAF'
	// Wing
	KahnWing = '\uEDE9'
	// Wing-change
	KahnWingChange = '\uEDEA'
	// Zank
	KahnZank = '\uEDE4'
	// Zink
	KahnZink = '\uEDDF'
	// Clavichord bebung, 2 finger movements (above)
	KeyboardBebung2DotsAbove = '\uE668'
	// Clavichord bebung, 2 finger movements (below)
	KeyboardBebung2DotsBelow = '\uE669'
	// Clavichord bebung, 3 finger movements (above)
	KeyboardBebung3DotsAbove = '\uE66A'
	// Clavichord bebung, 3 finger movements (below)
	KeyboardBebung3DotsBelow = '\uE66B'
	// Clavichord bebung, 4 finger movements (above)
	KeyboardBebung4DotsAbove = '\uE66C'
	// Clavichord bebung, 4 finger movements (below)
	KeyboardBebung4DotsBelow = '\uE66D'
	// Left pedal pictogram
	KeyboardLeftPedalPictogram = '\uE65E'
	// Middle pedal pictogram
	KeyboardMiddlePedalPictogram = '\uE65F'
	// Pedal d
	KeyboardPedalD = '\uE653'
	// Pedal dot
	KeyboardPedalDot = '\uE654'
	// Pedal e
	KeyboardPedalE = '\uE652'
	// Half-pedal mark
	KeyboardPedalHalf = '\uE656'
	// Half pedal mark 1
	KeyboardPedalHalf2 = '\uE65B'
	// Half pedal mark 2
	KeyboardPedalHalf3 = '\uE65C'
	// Pedal heel 1
	KeyboardPedalHeel1 = '\uE661'
	// Pedal heel 2
	KeyboardPedalHeel2 = '\uE662'
	// Pedal heel 3 (Davis)
	KeyboardPedalHeel3 = '\uE663'
	// Pedal heel to toe
	KeyboardPedalHeelToToe = '\uE674'
	// Pedal heel or toe
	KeyboardPedalHeelToe = '\uE666'
	// Pedal hook end
	KeyboardPedalHookEnd = '\uE673'
	// Pedal hook start
	KeyboardPedalHookStart = '\uE672'
	// Pedal hyphen
	KeyboardPedalHyphen = '\uE658'
	// Pedal P
	KeyboardPedalP = '\uE651'
	// Left parenthesis for pedal marking
	KeyboardPedalParensLeft = '\uE676'
	// Right parenthesis for pedal marking
	KeyboardPedalParensRight = '\uE677'
	// Pedal mark
	KeyboardPedalPed = '\uE650'
	// Pedal S
	KeyboardPedalS = '\uE65A'
	// Sostenuto pedal mark
	KeyboardPedalSost = '\uE659'
	// Pedal toe 1
	KeyboardPedalToe1 = '\uE664'
	// Pedal toe 2
	KeyboardPedalToe2 = '\uE665'
	// Pedal toe to heel
	KeyboardPedalToeToHeel = '\uE675'
	// Pedal up mark
	KeyboardPedalUp = '\uE655'
	// Pedal up notch
	KeyboardPedalUpNotch = '\uE657'
	// Pedal up special
	KeyboardPedalUpSpecial = '\uE65D'
	// Play with left hand
	KeyboardPlayWithLH = '\uE670'
	// Play with left hand (end)
	KeyboardPlayWithLHEnd = '\uE671'
	// Play with right hand
	KeyboardPlayWithRH = '\uE66E'
	// Play with right hand (end)
	KeyboardPlayWithRHEnd = '\uE66F'
	// Pluck strings inside piano (Maderna)
	KeyboardPluckInside = '\uE667'
	// Right pedal pictogram
	KeyboardRightPedalPictogram = '\uE660'
	// Kievan flat
	KievanAccidentalFlat = '\uEC3E'
	// Kievan sharp
	KievanAccidentalSharp = '\uEC3D'
	// Kievan augmentation dot
	KievanAugmentationDot = '\uEC3C'
	// Kievan C clef (tse-fa-ut)
	KievanCClef = '\uEC30'
	// Kievan ending symbol
	KievanEndingSymbol = '\uEC31'
	// Kievan eighth note, stem down
	KievanNote8thStemDown = '\uEC3A'
	// Kievan eighth note, stem up
	KievanNote8thStemUp = '\uEC39'
	// Kievan beam
	KievanNoteBeam = '\uEC3B'
	// Kievan half note (on staff line)
	KievanNoteHalfStaffLine = '\uEC35'
	// Kievan half note (in staff space)
	KievanNoteHalfStaffSpace = '\uEC36'
	// Kievan quarter note, stem down
	KievanNoteQuarterStemDown = '\uEC38'
	// Kievan quarter note, stem up
	KievanNoteQuarterStemUp = '\uEC37'
	// Kievan reciting note
	KievanNoteReciting = '\uEC32'
	// Kievan whole note
	KievanNoteWhole = '\uEC33'
	// Kievan final whole note
	KievanNoteWholeFinal = '\uEC34'
	// Do hand sign
	KodalyHandDo = '\uEC40'
	// Fa hand sign
	KodalyHandFa = '\uEC43'
	// La hand sign
	KodalyHandLa = '\uEC45'
	// Mi hand sign
	KodalyHandMi = '\uEC42'
	// Re hand sign
	KodalyHandRe = '\uEC41'
	// So hand sign
	KodalyHandSo = '\uEC44'
	// Ti hand sign
	KodalyHandTi = '\uEC46'
	// Left repeat sign within bar
	LeftRepeatSmall = '\uE04C'
	// Leger line
	LegerLine = '\uE022'
	// Leger line (narrow)
	LegerLineNarrow = '\uE024'
	// Leger line (wide)
	LegerLineWide = '\uE023'
	// Lute tablature end repeat barline
	LuteBarlineEndRepeat = '\uEBA4'
	// Lute tablature final barline
	LuteBarlineFinal = '\uEBA5'
	// Lute tablature start repeat barline
	LuteBarlineStartRepeat = '\uEBA3'
	// 16th note (semiquaver) duration sign
	LuteDuration16th = '\uEBAB'
	// 32nd note (demisemiquaver) duration sign
	LuteDuration32nd = '\uEBAC'
	// Eighth note (quaver) duration sign
	LuteDuration8th = '\uEBAA'
	// Double whole note (breve) duration sign
	LuteDurationDoubleWhole = '\uEBA6'
	// Half note (minim) duration sign
	LuteDurationHalf = '\uEBA8'
	// Quarter note (crotchet) duration sign
	LuteDurationQuarter = '\uEBA9'
	// Whole note (semibreve) duration sign
	LuteDurationWhole = '\uEBA7'
	// Right-hand fingering, first finger
	LuteFingeringRHFirst = '\uEBAE'
	// Right-hand fingering, second finger
	LuteFingeringRHSecond = '\uEBAF'
	// Right-hand fingering, third finger
	LuteFingeringRHThird = '\uEBB0'
	// Right-hand fingering, thumb
	LuteFingeringRHThumb = '\uEBAD'
	// 10th course (diapason)
	LuteFrench10thCourse = '\uEBD0'
	// Seventh course (diapason)
	LuteFrench7thCourse = '\uEBCD'
	// Eighth course (diapason)
	LuteFrench8thCourse = '\uEBCE'
	// Ninth course (diapason)
	LuteFrench9thCourse = '\uEBCF'
	// Appoggiatura from above
	LuteFrenchAppoggiaturaAbove = '\uEBD5'
	// Appoggiatura from below
	LuteFrenchAppoggiaturaBelow = '\uEBD4'
	// Open string (a)
	LuteFrenchFretA = '\uEBC0'
	// First fret (b)
	LuteFrenchFretB = '\uEBC1'
	// Second fret (c)
	LuteFrenchFretC = '\uEBC2'
	// Third fret (d)
	LuteFrenchFretD = '\uEBC3'
	// Fourth fret (e)
	LuteFrenchFretE = '\uEBC4'
	// Fifth fret (f)
	LuteFrenchFretF = '\uEBC5'
	// Sixth fret (g)
	LuteFrenchFretG = '\uEBC6'
	// Seventh fret (h)
	LuteFrenchFretH = '\uEBC7'
	// Eighth fret (i)
	LuteFrenchFretI = '\uEBC8'
	// Ninth fret (k)
	LuteFrenchFretK = '\uEBC9'
	// 10th fret (l)
	LuteFrenchFretL = '\uEBCA'
	// 11th fret (m)
	LuteFrenchFretM = '\uEBCB'
	// 12th fret (n)
	LuteFrenchFretN = '\uEBCC'
	// Inverted mordent
	LuteFrenchMordentInverted = '\uEBD3'
	// Mordent with lower auxiliary
	LuteFrenchMordentLower = '\uEBD2'
	// Mordent with upper auxiliary
	LuteFrenchMordentUpper = '\uEBD1'
	// 5th course, 1st fret (a)
	LuteGermanALower = '\uEC00'
	// 6th course, 1st fret (A)
	LuteGermanAUpper = '\uEC17'
	// 4th course, 1st fret (b)
	LuteGermanBLower = '\uEC01'
	// 6th course, 2nd fret (B)
	LuteGermanBUpper = '\uEC18'
	// 3rd course, 1st fret (c)
	LuteGermanCLower = '\uEC02'
	// 6th course, 3rd fret (C)
	LuteGermanCUpper = '\uEC19'
	// 2nd course, 1st fret (d)
	LuteGermanDLower = '\uEC03'
	// 6th course, 4th fret (D)
	LuteGermanDUpper = '\uEC1A'
	// 1st course, 1st fret (e)
	LuteGermanELower = '\uEC04'
	// 6th course, 5th fret (E)
	LuteGermanEUpper = '\uEC1B'
	// 5th course, 2nd fret (f)
	LuteGermanFLower = '\uEC05'
	// 6th course, 6th fret (F)
	LuteGermanFUpper = '\uEC1C'
	// 4th course, 2nd fret (g)
	LuteGermanGLower = '\uEC06'
	// 6th course, 7th fret (G)
	LuteGermanGUpper = '\uEC1D'
	// 3rd course, 2nd fret (h)
	LuteGermanHLower = '\uEC07'
	// 6th course, 8th fret (H)
	LuteGermanHUpper = '\uEC1E'
	// 2nd course, 2nd fret (i)
	LuteGermanILower = '\uEC08'
	// 6th course, 9th fret (I)
	LuteGermanIUpper = '\uEC1F'
	// 1st course, 2nd fret (k)
	LuteGermanKLower = '\uEC09'
	// 6th course, 10th fret (K)
	LuteGermanKUpper = '\uEC20'
	// 5th course, 3rd fret (l)
	LuteGermanLLower = '\uEC0A'
	// 6th course, 11th fret (L)
	LuteGermanLUpper = '\uEC21'
	// 4th course, 3rd fret (m)
	LuteGermanMLower = '\uEC0B'
	// 6th course, 12th fret (M)
	LuteGermanMUpper = '\uEC22'
	// 3rd course, 3rd fret (n)
	LuteGermanNLower = '\uEC0C'
	// 6th course, 13th fret (N)
	LuteGermanNUpper = '\uEC23'
	// 2nd course, 3rd fret (o)
	LuteGermanOLower = '\uEC0D'
	// 1st course, 3rd fret (p)
	LuteGermanPLower = '\uEC0E'
	// 5th course, 4th fret (q)
	LuteGermanQLower = '\uEC0F'
	// 4th course, 4th fret (r)
	LuteGermanRLower = '\uEC10'
	// 3rd course, 4th fret (s)
	LuteGermanSLower = '\uEC11'
	// 2nd course, 4th fret (t)
	LuteGermanTLower = '\uEC12'
	// 1st course, 4th fret (v)
	LuteGermanVLower = '\uEC13'
	// 5th course, 5th fret (x)
	LuteGermanXLower = '\uEC14'
	// 4th course, 5th fret (y)
	LuteGermanYLower = '\uEC15'
	// 3rd course, 5th fret (z)
	LuteGermanZLower = '\uEC16'
	// C sol fa ut clef
	LuteItalianClefCSolFaUt = '\uEBF1'
	// F fa ut clef
	LuteItalianClefFFaUt = '\uEBF0'
	// Open string (0)
	LuteItalianFret0 = '\uEBE0'
	// First fret (1)
	LuteItalianFret1 = '\uEBE1'
	// Second fret (2)
	LuteItalianFret2 = '\uEBE2'
	// Third fret (3)
	LuteItalianFret3 = '\uEBE3'
	// Fourth fret (4)
	LuteItalianFret4 = '\uEBE4'
	// Fifth fret (5)
	LuteItalianFret5 = '\uEBE5'
	// Sixth fret (6)
	LuteItalianFret6 = '\uEBE6'
	// Seventh fret (7)
	LuteItalianFret7 = '\uEBE7'
	// Eighth fret (8)
	LuteItalianFret8 = '\uEBE8'
	// Ninth fret (9)
	LuteItalianFret9 = '\uEBE9'
	// Hold finger in place
	LuteItalianHoldFinger = '\uEBF4'
	// Hold note
	LuteItalianHoldNote = '\uEBF3'
	// Release finger
	LuteItalianReleaseFinger = '\uEBF5'
	// Fast tempo indication (de Mudarra)
	LuteItalianTempoFast = '\uEBEA'
	// Neither fast nor slow tempo indication (de Mudarra)
	LuteItalianTempoNeitherFastNorSlow = '\uEBEC'
	// Slow tempo indication (de Mudarra)
	LuteItalianTempoSlow = '\uEBED'
	// Somewhat fast tempo indication (de Narvaez)
	LuteItalianTempoSomewhatFast = '\uEBEB'
	// Very slow indication (de Narvaez)
	LuteItalianTempoVerySlow = '\uEBEE'
	// Triple time indication
	LuteItalianTimeTriple = '\uEBEF'
	// Single-finger tremolo or mordent
	LuteItalianTremolo = '\uEBF2'
	// Vibrato (verre cassé)
	LuteItalianVibrato = '\uEBF6'
	// Lute tablature staff, 6 courses
	LuteStaff6Lines = '\uEBA0'
	// Lute tablature staff, 6 courses (narrow)
	LuteStaff6LinesNarrow = '\uEBA2'
	// Lute tablature staff, 6 courses (wide)
	LuteStaff6LinesWide = '\uEBA1'
	// Elision
	LyricsElision = '\uE551'
	// Narrow elision
	LyricsElisionNarrow = '\uE550'
	// Wide elision
	LyricsElisionWide = '\uE552'
	// Baseline hyphen
	LyricsHyphenBaseline = '\uE553'
	// Non-breaking baseline hyphen
	LyricsHyphenBaselineNonBreaking = '\uE554'
	// Text repeats
	LyricsTextRepeat = '\uE555'
	// Flat, hard b (mi)
	MedRenFlatHardB = '\uE9E1'
	// Flat, soft b (fa)
	MedRenFlatSoftB = '\uE9E0'
	// Flat with dot
	MedRenFlatWithDot = '\uE9E4'
	// G clef (Corpus Monodicum)
	MedRenGClefCMN = '\uEA24'
	// Liquescence
	MedRenLiquescenceCMN = '\uEA22'
	// Liquescent ascending (Corpus Monodicum)
	MedRenLiquescentAscCMN = '\uEA26'
	// Liquescent descending (Corpus Monodicum)
	MedRenLiquescentDescCMN = '\uEA27'
	// Natural
	MedRenNatural = '\uE9E2'
	// Natural with interrupted cross
	MedRenNaturalWithCross = '\uE9E5'
	// Oriscus (Corpus Monodicum)
	MedRenOriscusCMN = '\uEA2A'
	// Plica
	MedRenPlicaCMN = '\uEA23'
	// Punctum (Corpus Monodicum)
	MedRenPunctumCMN = '\uEA25'
	// Quilisma (Corpus Monodicum)
	MedRenQuilismaCMN = '\uEA28'
	// Croix
	MedRenSharpCroix = '\uE9E3'
	// Strophicus (Corpus Monodicum)
	MedRenStrophicusCMN = '\uEA29'
	// Alteration sign
	MensuralAlterationSign = '\uEA10'
	// Black mensural brevis
	MensuralBlackBrevis = '\uE952'
	// Black mensural void brevis
	MensuralBlackBrevisVoid = '\uE956'
	// Black mensural dragma
	MensuralBlackDragma = '\uE95A'
	// Black mensural longa
	MensuralBlackLonga = '\uE951'
	// Black mensural maxima
	MensuralBlackMaxima = '\uE950'
	// Black mensural minima
	MensuralBlackMinima = '\uE954'
	// Black mensural void minima
	MensuralBlackMinimaVoid = '\uE958'
	// Black mensural semibrevis
	MensuralBlackSemibrevis = '\uE953'
	// Black mensural semibrevis caudata
	MensuralBlackSemibrevisCaudata = '\uE959'
	// Black mensural oblique semibrevis
	MensuralBlackSemibrevisOblique = '\uE95B'
	// Black mensural void semibrevis
	MensuralBlackSemibrevisVoid = '\uE957'
	// Black mensural semiminima
	MensuralBlackSemiminima = '\uE955'
	// Mensural C clef
	MensuralCclef = '\uE905'
	// Petrucci C clef, high position
	MensuralCclefPetrucciPosHigh = '\uE90A'
	// Petrucci C clef, highest position
	MensuralCclefPetrucciPosHighest = '\uE90B'
	// Petrucci C clef, low position
	MensuralCclefPetrucciPosLow = '\uE908'
	// Petrucci C clef, lowest position
	MensuralCclefPetrucciPosLowest = '\uE907'
	// Petrucci C clef, middle position
	MensuralCclefPetrucciPosMiddle = '\uE909'
	// Coloration end, round
	MensuralColorationEndRound = '\uEA0F'
	// Coloration end, square
	MensuralColorationEndSquare = '\uEA0D'
	// Coloration start, round
	MensuralColorationStartRound = '\uEA0E'
	// Coloration start, square
	MensuralColorationStartSquare = '\uEA0C'
	// Combining stem diagonal
	MensuralCombStemDiagonal = '\uE940'
	// Combining stem down
	MensuralCombStemDown = '\uE93F'
	// Combining stem with extended flag down
	MensuralCombStemDownFlagExtended = '\uE948'
	// Combining stem with flared flag down
	MensuralCombStemDownFlagFlared = '\uE946'
	// Combining stem with fusa flag down
	MensuralCombStemDownFlagFusa = '\uE94C'
	// Combining stem with flag left down
	MensuralCombStemDownFlagLeft = '\uE944'
	// Combining stem with flag right down
	MensuralCombStemDownFlagRight = '\uE942'
	// Combining stem with semiminima flag down
	MensuralCombStemDownFlagSemiminima = '\uE94A'
	// Combining stem up
	MensuralCombStemUp = '\uE93E'
	// Combining stem with extended flag up
	MensuralCombStemUpFlagExtended = '\uE947'
	// Combining stem with flared flag up
	MensuralCombStemUpFlagFlared = '\uE945'
	// Combining stem with fusa flag up
	MensuralCombStemUpFlagFusa = '\uE94B'
	// Combining stem with flag left up
	MensuralCombStemUpFlagLeft = '\uE943'
	// Combining stem with flag right up
	MensuralCombStemUpFlagRight = '\uE941'
	// Combining stem with semiminima flag up
	MensuralCombStemUpFlagSemiminima = '\uE949'
	// Checkmark custos
	MensuralCustosCheckmark = '\uEA0A'
	// Mensural custos down
	MensuralCustosDown = '\uEA03'
	// Turn-like custos
	MensuralCustosTurn = '\uEA0B'
	// Mensural custos up
	MensuralCustosUp = '\uEA02'
	// Mensural F clef
	MensuralFclef = '\uE903'
	// Petrucci F clef
	MensuralFclefPetrucci = '\uE904'
	// Mensural G clef
	MensuralGclef = '\uE900'
	// Petrucci G clef
	MensuralGclefPetrucci = '\uE901'
	// Modus imperfectum, vertical
	MensuralModusImperfectumVert = '\uE92D'
	// Modus perfectum, vertical
	MensuralModusPerfectumVert = '\uE92C'
	// Longa/brevis notehead, black
	MensuralNoteheadLongaBlack = '\uE934'
	// Longa/brevis notehead, black and void
	MensuralNoteheadLongaBlackVoid = '\uE936'
	// Longa/brevis notehead, void
	MensuralNoteheadLongaVoid = '\uE935'
	// Longa/brevis notehead, white
	MensuralNoteheadLongaWhite = '\uE937'
	// Maxima notehead, black
	MensuralNoteheadMaximaBlack = '\uE930'
	// Maxima notehead, black and void
	MensuralNoteheadMaximaBlackVoid = '\uE932'
	// Maxima notehead, void
	MensuralNoteheadMaximaVoid = '\uE931'
	// Maxima notehead, white
	MensuralNoteheadMaximaWhite = '\uE933'
	// Minima notehead, white
	MensuralNoteheadMinimaWhite = '\uE93C'
	// Semibrevis notehead, black
	MensuralNoteheadSemibrevisBlack = '\uE938'
	// Semibrevis notehead, black and void
	MensuralNoteheadSemibrevisBlackVoid = '\uE93A'
	// Semibrevis notehead, black and void (turned)
	MensuralNoteheadSemibrevisBlackVoidTurned = '\uE93B'
	// Semibrevis notehead, void
	MensuralNoteheadSemibrevisVoid = '\uE939'
	// Semiminima/fusa notehead, white
	MensuralNoteheadSemiminimaWhite = '\uE93D'
	// Oblique form, ascending 2nd, black
	MensuralObliqueAsc2ndBlack = '\uE970'
	// Oblique form, ascending 2nd, black and void
	MensuralObliqueAsc2ndBlackVoid = '\uE972'
	// Oblique form, ascending 2nd, void
	MensuralObliqueAsc2ndVoid = '\uE971'
	// Oblique form, ascending 2nd, white
	MensuralObliqueAsc2ndWhite = '\uE973'
	// Oblique form, ascending 3rd, black
	MensuralObliqueAsc3rdBlack = '\uE974'
	// Oblique form, ascending 3rd, black and void
	MensuralObliqueAsc3rdBlackVoid = '\uE976'
	// Oblique form, ascending 3rd, void
	MensuralObliqueAsc3rdVoid = '\uE975'
	// Oblique form, ascending 3rd, white
	MensuralObliqueAsc3rdWhite = '\uE977'
	// Oblique form, ascending 4th, black
	MensuralObliqueAsc4thBlack = '\uE978'
	// Oblique form, ascending 4th, black and void
	MensuralObliqueAsc4thBlackVoid = '\uE97A'
	// Oblique form, ascending 4th, void
	MensuralObliqueAsc4thVoid = '\uE979'
	// Oblique form, ascending 4th, white
	MensuralObliqueAsc4thWhite = '\uE97B'
	// Oblique form, ascending 5th, black
	MensuralObliqueAsc5thBlack = '\uE97C'
	// Oblique form, ascending 5th, black and void
	MensuralObliqueAsc5thBlackVoid = '\uE97E'
	// Oblique form, ascending 5th, void
	MensuralObliqueAsc5thVoid = '\uE97D'
	// Oblique form, ascending 5th, white
	MensuralObliqueAsc5thWhite = '\uE97F'
	// Oblique form, descending 2nd, black
	MensuralObliqueDesc2ndBlack = '\uE980'
	// Oblique form, descending 2nd, black and void
	MensuralObliqueDesc2ndBlackVoid = '\uE982'
	// Oblique form, descending 2nd, void
	MensuralObliqueDesc2ndVoid = '\uE981'
	// Oblique form, descending 2nd, white
	MensuralObliqueDesc2ndWhite = '\uE983'
	// Oblique form, descending 3rd, black
	MensuralObliqueDesc3rdBlack = '\uE984'
	// Oblique form, descending 3rd, black and void
	MensuralObliqueDesc3rdBlackVoid = '\uE986'
	// Oblique form, descending 3rd, void
	MensuralObliqueDesc3rdVoid = '\uE985'
	// Oblique form, descending 3rd, white
	MensuralObliqueDesc3rdWhite = '\uE987'
	// Oblique form, descending 4th, black
	MensuralObliqueDesc4thBlack = '\uE988'
	// Oblique form, descending 4th, black and void
	MensuralObliqueDesc4thBlackVoid = '\uE98A'
	// Oblique form, descending 4th, void
	MensuralObliqueDesc4thVoid = '\uE989'
	// Oblique form, descending 4th, white
	MensuralObliqueDesc4thWhite = '\uE98B'
	// Oblique form, descending 5th, black
	MensuralObliqueDesc5thBlack = '\uE98C'
	// Oblique form, descending 5th, black and void
	MensuralObliqueDesc5thBlackVoid = '\uE98E'
	// Oblique form, descending 5th, void
	MensuralObliqueDesc5thVoid = '\uE98D'
	// Oblique form, descending 5th, white
	MensuralObliqueDesc5thWhite = '\uE98F'
	// Tempus perfectum cum prolatione perfecta (9/8)
	MensuralProlation1 = '\uE910'
	// Tempus imperfectum cum prolatione imperfecta diminution 4
	MensuralProlation10 = '\uE919'
	// Tempus imperfectum cum prolatione imperfecta diminution 5
	MensuralProlation11 = '\uE91A'
	// Tempus perfectum cum prolatione imperfecta (3/4)
	MensuralProlation2 = '\uE911'
	// Tempus perfectum cum prolatione imperfecta diminution 1 (3/8)
	MensuralProlation3 = '\uE912'
	// Tempus perfectum cum prolatione perfecta diminution 2 (9/16)
	MensuralProlation4 = '\uE913'
	// Tempus imperfectum cum prolatione perfecta (6/8)
	MensuralProlation5 = '\uE914'
	// Tempus imperfectum cum prolatione imperfecta (2/4)
	MensuralProlation6 = '\uE915'
	// Tempus imperfectum cum prolatione imperfecta diminution 1 (2/2)
	MensuralProlation7 = '\uE916'
	// Tempus imperfectum cum prolatione imperfecta diminution 2 (6/16)
	MensuralProlation8 = '\uE917'
	// Tempus imperfectum cum prolatione imperfecta diminution 3 (2/2)
	MensuralProlation9 = '\uE918'
	// Combining dot
	MensuralProlationCombiningDot = '\uE920'
	// Combining void dot
	MensuralProlationCombiningDotVoid = '\uE924'
	// Combining vertical stroke
	MensuralProlationCombiningStroke = '\uE925'
	// Combining three dots horizontal
	MensuralProlationCombiningThreeDots = '\uE922'
	// Combining three dots triangular
	MensuralProlationCombiningThreeDotsTri = '\uE923'
	// Combining two dots
	MensuralProlationCombiningTwoDots = '\uE921'
	// Mensural proportion 1
	MensuralProportion1 = '\uE926'
	// Mensural proportion 2
	MensuralProportion2 = '\uE927'
	// Mensural proportion 3
	MensuralProportion3 = '\uE928'
	// Mensural proportion 4
	MensuralProportion4 = '\uE929'
	// Mensural proportion 5
	MensuralProportion5 = '\uEE90'
	// Mensural proportion 6
	MensuralProportion6 = '\uEE91'
	// Mensural proportion 7
	MensuralProportion7 = '\uEE92'
	// Mensural proportion 8
	MensuralProportion8 = '\uEE93'
	// Mensural proportion 9
	MensuralProportion9 = '\uEE94'
	// Mensural proportion major
	MensuralProportionMajor = '\uE92B'
	// Mensural proportion minor
	MensuralProportionMinor = '\uE92A'
	// Proportio dupla 1
	MensuralProportionProportioDupla1 = '\uE91C'
	// Proportio dupla 2
	MensuralProportionProportioDupla2 = '\uE91D'
	// Proportio quadrupla
	MensuralProportionProportioQuadrupla = '\uE91F'
	// Proportio tripla
	MensuralProportionProportioTripla = '\uE91E'
	// Tempus perfectum
	MensuralProportionTempusPerfectum = '\uE91B'
	// Brevis rest
	MensuralRestBrevis = '\uE9F3'
	// Fusa rest
	MensuralRestFusa = '\uE9F7'
	// Longa imperfecta rest
	MensuralRestLongaImperfecta = '\uE9F2'
	// Longa perfecta rest
	MensuralRestLongaPerfecta = '\uE9F1'
	// Maxima rest
	MensuralRestMaxima = '\uE9F0'
	// Minima rest
	MensuralRestMinima = '\uE9F5'
	// Semibrevis rest
	MensuralRestSemibrevis = '\uE9F4'
	// Semifusa rest
	MensuralRestSemifusa = '\uE9F8'
	// Semiminima rest
	MensuralRestSemiminima = '\uE9F6'
	// Signum congruentiae down
	MensuralSignumDown = '\uEA01'
	// Signum congruentiae up
	MensuralSignumUp = '\uEA00'
	// Tempus imperfectum, horizontal
	MensuralTempusImperfectumHoriz = '\uE92F'
	// Tempus perfectum, horizontal
	MensuralTempusPerfectumHoriz = '\uE92E'
	// White mensural brevis
	MensuralWhiteBrevis = '\uE95E'
	// White mensural fusa
	MensuralWhiteFusa = '\uE961'
	// White mensural longa
	MensuralWhiteLonga = '\uE95D'
	// White mensural maxima
	MensuralWhiteMaxima = '\uE95C'
	// White mensural minima
	MensuralWhiteMinima = '\uE95F'
	// White mensural semibrevis
	MensuralWhiteSemibrevis = '\uE962'
	// White mensural semiminima
	MensuralWhiteSemiminima = '\uE960'
	// Augmentation dot
	MetAugmentationDot = '\uECB7'
	// 1024th note (semihemidemisemihemidemisemiquaver) stem down
	MetNote1024thDown = '\uECB6'
	// 1024th note (semihemidemisemihemidemisemiquaver) stem up
	MetNote1024thUp = '\uECB5'
	// 128th note (semihemidemisemiquaver) stem down
	MetNote128thDown = '\uECB0'
	// 128th note (semihemidemisemiquaver) stem up
	MetNote128thUp = '\uECAF'
	// 16th note (semiquaver) stem down
	MetNote16thDown = '\uECAA'
	// 16th note (semiquaver) stem up
	MetNote16thUp = '\uECA9'
	// 256th note (demisemihemidemisemiquaver) stem down
	MetNote256thDown = '\uECB2'
	// 256th note (demisemihemidemisemiquaver) stem up
	MetNote256thUp = '\uECB1'
	// 32nd note (demisemiquaver) stem down
	MetNote32ndDown = '\uECAC'
	// 32nd note (demisemiquaver) stem up
	MetNote32ndUp = '\uECAB'
	// 512th note (hemidemisemihemidemisemiquaver) stem down
	MetNote512thDown = '\uECB4'
	// 512th note (hemidemisemihemidemisemiquaver) stem up
	MetNote512thUp = '\uECB3'
	// 64th note (hemidemisemiquaver) stem down
	MetNote64thDown = '\uECAE'
	// 64th note (hemidemisemiquaver) stem up
	MetNote64thUp = '\uECAD'
	// Eighth note (quaver) stem down
	MetNote8thDown = '\uECA8'
	// Eighth note (quaver) stem up
	MetNote8thUp = '\uECA7'
	// Double whole note (breve)
	MetNoteDoubleWhole = '\uECA0'
	// Double whole note (square)
	MetNoteDoubleWholeSquare = '\uECA1'
	// Half note (minim) stem down
	MetNoteHalfDown = '\uECA4'
	// Half note (minim) stem up
	MetNoteHalfUp = '\uECA3'
	// Quarter note (crotchet) stem down
	MetNoteQuarterDown = '\uECA6'
	// Quarter note (crotchet) stem up
	MetNoteQuarterUp = '\uECA5'
	// Whole note (semibreve)
	MetNoteWhole = '\uECA2'
	// Left-pointing arrow for metric modulation
	MetricModulationArrowLeft = '\uEC63'
	// Right-pointing arrow for metric modulation
	MetricModulationArrowRight = '\uEC64'
	// Do not copy
	MiscDoNotCopy = '\uEC61'
	// Do not photocopy
	MiscDoNotPhotocopy = '\uEC60'
	// Eyeglasses
	MiscEyeglasses = '\uEC62'
	// 1024th note (semihemidemisemihemidemisemiquaver) stem down
	Note1024thDown = '\uE1E6'
	// 1024th note (semihemidemisemihemidemisemiquaver) stem up
	Note1024thUp = '\uE1E5'
	// 128th note (semihemidemisemiquaver) stem down
	Note128thDown = '\uE1E0'
	// 128th note (semihemidemisemiquaver) stem up
	Note128thUp = '\uE1DF'
	// 16th note (semiquaver) stem down
	Note16thDown = '\uE1DA'
	// 16th note (semiquaver) stem up
	Note16thUp = '\uE1D9'
	// 256th note (demisemihemidemisemiquaver) stem down
	Note256thDown = '\uE1E2'
	// 256th note (demisemihemidemisemiquaver) stem up
	Note256thUp = '\uE1E1'
	// 32nd note (demisemiquaver) stem down
	Note32ndDown = '\uE1DC'
	// 32nd note (demisemiquaver) stem up
	Note32ndUp = '\uE1DB'
	// 512th note (hemidemisemihemidemisemiquaver) stem down
	Note512thDown = '\uE1E4'
	// 512th note (hemidemisemihemidemisemiquaver) stem up
	Note512thUp = '\uE1E3'
	// 64th note (hemidemisemiquaver) stem down
	Note64thDown = '\uE1DE'
	// 64th note (hemidemisemiquaver) stem up
	Note64thUp = '\uE1DD'
	// Eighth note (quaver) stem down
	Note8thDown = '\uE1D8'
	// Eighth note (quaver) stem up
	Note8thUp = '\uE1D7'
	// A (black note)
	NoteABlack = '\uE197'
	// A flat (black note)
	NoteAFlatBlack = '\uE196'
	// A flat (half note)
	NoteAFlatHalf = '\uE17F'
	// A flat (whole note)
	NoteAFlatWhole = '\uE168'
	// A (half note)
	NoteAHalf = '\uE180'
	// A sharp (black note)
	NoteASharpBlack = '\uE198'
	// A sharp (half note)
	NoteASharpHalf = '\uE181'
	// A sharp (whole note)
	NoteASharpWhole = '\uE16A'
	// A (whole note)
	NoteAWhole = '\uE169'
	// B (black note)
	NoteBBlack = '\uE19A'
	// B flat (black note)
	NoteBFlatBlack = '\uE199'
	// B flat (half note)
	NoteBFlatHalf = '\uE182'
	// B flat (whole note)
	NoteBFlatWhole = '\uE16B'
	// B (half note)
	NoteBHalf = '\uE183'
	// B sharp (black note)
	NoteBSharpBlack = '\uE19B'
	// B sharp (half note)
	NoteBSharpHalf = '\uE184'
	// B sharp (whole note)
	NoteBSharpWhole = '\uE16D'
	// B (whole note)
	NoteBWhole = '\uE16C'
	// C (black note)
	NoteCBlack = '\uE19D'
	// C flat (black note)
	NoteCFlatBlack = '\uE19C'
	// C flat (half note)
	NoteCFlatHalf = '\uE185'
	// C flat (whole note)
	NoteCFlatWhole = '\uE16E'
	// C (half note)
	NoteCHalf = '\uE186'
	// C sharp (black note)
	NoteCSharpBlack = '\uE19E'
	// C sharp (half note)
	NoteCSharpHalf = '\uE187'
	// C sharp (whole note)
	NoteCSharpWhole = '\uE170'
	// C (whole note)
	NoteCWhole = '\uE16F'
	// D (black note)
	NoteDBlack = '\uE1A0'
	// D flat (black note)
	NoteDFlatBlack = '\uE19F'
	// D flat (half note)
	NoteDFlatHalf = '\uE188'
	// D flat (whole note)
	NoteDFlatWhole = '\uE171'
	// D (half note)
	NoteDHalf = '\uE189'
	// D sharp (black note)
	NoteDSharpBlack = '\uE1A1'
	// D sharp (half note)
	NoteDSharpHalf = '\uE18A'
	// D sharp (whole note)
	NoteDSharpWhole = '\uE173'
	// D (whole note)
	NoteDWhole = '\uE172'
	// Di (black note)
	NoteDiBlack = '\uEEF2'
	// Di (half note)
	NoteDiHalf = '\uEEE9'
	// Di (whole note)
	NoteDiWhole = '\uEEE0'
	// Do (black note)
	NoteDoBlack = '\uE160'
	// Do (half note)
	NoteDoHalf = '\uE158'
	// Do (whole note)
	NoteDoWhole = '\uE150'
	// Double whole note (breve)
	NoteDoubleWhole = '\uE1D0'
	// Double whole note (square)
	NoteDoubleWholeSquare = '\uE1D1'
	// E (black note)
	NoteEBlack = '\uE1A3'
	// E flat (black note)
	NoteEFlatBlack = '\uE1A2'
	// E flat (half note)
	NoteEFlatHalf = '\uE18B'
	// E flat (whole note)
	NoteEFlatWhole = '\uE174'
	// E (half note)
	NoteEHalf = '\uE18C'
	// E sharp (black note)
	NoteESharpBlack = '\uE1A4'
	// E sharp (half note)
	NoteESharpHalf = '\uE18D'
	// E sharp (whole note)
	NoteESharpWhole = '\uE176'
	// E (whole note)
	NoteEWhole = '\uE175'
	// Empty black note
	NoteEmptyBlack = '\uE1AF'
	// Empty half note
	NoteEmptyHalf = '\uE1AE'
	// Empty whole note
	NoteEmptyWhole = '\uE1AD'
	// F (black note)
	NoteFBlack = '\uE1A6'
	// F flat (black note)
	NoteFFlatBlack = '\uE1A5'
	// F flat (half note)
	NoteFFlatHalf = '\uE18E'
	// F flat (whole note)
	NoteFFlatWhole = '\uE177'
	// F (half note)
	NoteFHalf = '\uE18F'
	// F sharp (black note)
	NoteFSharpBlack = '\uE1A7'
	// F sharp (half note)
	NoteFSharpHalf = '\uE190'
	// F sharp (whole note)
	NoteFSharpWhole = '\uE179'
	// F (whole note)
	NoteFWhole = '\uE178'
	// Fa (black note)
	NoteFaBlack = '\uE163'
	// Fa (half note)
	NoteFaHalf = '\uE15B'
	// Fa (whole note)
	NoteFaWhole = '\uE153'
	// Fi (black note)
	NoteFiBlack = '\uEEF6'
	// Fi (half note)
	NoteFiHalf = '\uEEED'
	// Fi (whole note)
	NoteFiWhole = '\uEEE4'
	// G (black note)
	NoteGBlack = '\uE1A9'
	// G flat (black note)
	NoteGFlatBlack = '\uE1A8'
	// G flat (half note)
	NoteGFlatHalf = '\uE191'
	// G flat (whole note)
	NoteGFlatWhole = '\uE17A'
	// G (half note)
	NoteGHalf = '\uE192'
	// G sharp (black note)
	NoteGSharpBlack = '\uE1AA'
	// G sharp (half note)
	NoteGSharpHalf = '\uE193'
	// G sharp (whole note)
	NoteGSharpWhole = '\uE17C'
	// G (whole note)
	NoteGWhole = '\uE17B'
	// H (black note)
	NoteHBlack = '\uE1AB'
	// H (half note)
	NoteHHalf = '\uE194'
	// H sharp (black note)
	NoteHSharpBlack = '\uE1AC'
	// H sharp (half note)
	NoteHSharpHalf = '\uE195'
	// H sharp (whole note)
	NoteHSharpWhole = '\uE17E'
	// H (whole note)
	NoteHWhole = '\uE17D'
	// Half note (minim) stem down
	NoteHalfDown = '\uE1D4'
	// Half note (minim) stem up
	NoteHalfUp = '\uE1D3'
	// La (black note)
	NoteLaBlack = '\uE165'
	// La (half note)
	NoteLaHalf = '\uE15D'
	// La (whole note)
	NoteLaWhole = '\uE155'
	// Le (black note)
	NoteLeBlack = '\uEEF9'
	// Le (half note)
	NoteLeHalf = '\uEEF0'
	// Le (whole note)
	NoteLeWhole = '\uEEE7'
	// Li (black note)
	NoteLiBlack = '\uEEF8'
	// Li (half note)
	NoteLiHalf = '\uEEEF'
	// Li (whole note)
	NoteLiWhole = '\uEEE6'
	// Me (black note)
	NoteMeBlack = '\uEEF5'
	// Me (half note)
	NoteMeHalf = '\uEEEC'
	// Me (whole note)
	NoteMeWhole = '\uEEE3'
	// Mi (black note)
	NoteMiBlack = '\uE162'
	// Mi (half note)
	NoteMiHalf = '\uE15A'
	// Mi (whole note)
	NoteMiWhole = '\uE152'
	// Quarter note (crotchet) stem down
	NoteQuarterDown = '\uE1D6'
	// Quarter note (crotchet) stem up
	NoteQuarterUp = '\uE1D5'
	// Ra (black note)
	NoteRaBlack = '\uEEF4'
	// Ra (half note)
	NoteRaHalf = '\uEEEB'
	// Ra (whole note)
	NoteRaWhole = '\uEEE2'
	// Re (black note)
	NoteReBlack = '\uE161'
	// Re (half note)
	NoteReHalf = '\uE159'
	// Re (whole note)
	NoteReWhole = '\uE151'
	// Ri (black note)
	NoteRiBlack = '\uEEF3'
	// Ri (half note)
	NoteRiHalf = '\uEEEA'
	// Ri (whole note)
	NoteRiWhole = '\uEEE1'
	// Se (black note)
	NoteSeBlack = '\uEEF7'
	// Se (half note)
	NoteSeHalf = '\uEEEE'
	// Se (whole note)
	NoteSeWhole = '\uEEE5'
	// Arrowhead left black (Funk 7-shape re)
	NoteShapeArrowheadLeftBlack = '\uE1C9'
	// Arrowhead left double whole (Funk 7-shape re)
	NoteShapeArrowheadLeftDoubleWhole = '\uECDC'
	// Arrowhead left white (Funk 7-shape re)
	NoteShapeArrowheadLeftWhite = '\uE1C8'
	// Diamond black (4-shape mi; 7-shape mi)
	NoteShapeDiamondBlack = '\uE1B9'
	// Diamond double whole (4-shape mi; 7-shape mi)
	NoteShapeDiamondDoubleWhole = '\uECD4'
	// Diamond white (4-shape mi; 7-shape mi)
	NoteShapeDiamondWhite = '\uE1B8'
	// Isosceles triangle black (Walker 7-shape ti)
	NoteShapeIsoscelesTriangleBlack = '\uE1C5'
	// Isosceles triangle double whole (Walker 7-shape ti)
	NoteShapeIsoscelesTriangleDoubleWhole = '\uECDA'
	// Isosceles triangle white (Walker 7-shape ti)
	NoteShapeIsoscelesTriangleWhite = '\uE1C4'
	// Inverted keystone black (Walker 7-shape do)
	NoteShapeKeystoneBlack = '\uE1C1'
	// Inverted keystone double whole (Walker 7-shape do)
	NoteShapeKeystoneDoubleWhole = '\uECD8'
	// Inverted keystone white (Walker 7-shape do)
	NoteShapeKeystoneWhite = '\uE1C0'
	// Moon black (Aikin 7-shape re)
	NoteShapeMoonBlack = '\uE1BD'
	// Moon double whole (Aikin 7-shape re)
	NoteShapeMoonDoubleWhole = '\uECD6'
	// Moon left black (Funk 7-shape do)
	NoteShapeMoonLeftBlack = '\uE1C7'
	// Moon left double whole (Funk 7-shape do)
	NoteShapeMoonLeftDoubleWhole = '\uECDB'
	// Moon left white (Funk 7-shape do)
	NoteShapeMoonLeftWhite = '\uE1C6'
	// Moon white (Aikin 7-shape re)
	NoteShapeMoonWhite = '\uE1BC'
	// Quarter moon black (Walker 7-shape re)
	NoteShapeQuarterMoonBlack = '\uE1C3'
	// Quarter moon double whole (Walker 7-shape re)
	NoteShapeQuarterMoonDoubleWhole = '\uECD9'
	// Quarter moon white (Walker 7-shape re)
	NoteShapeQuarterMoonWhite = '\uE1C2'
	// Round black (4-shape sol; 7-shape so)
	NoteShapeRoundBlack = '\uE1B1'
	// Round double whole (4-shape sol; 7-shape so)
	NoteShapeRoundDoubleWhole = '\uECD0'
	// Round white (4-shape sol; 7-shape so)
	NoteShapeRoundWhite = '\uE1B0'
	// Square black (4-shape la; Aikin 7-shape la)
	NoteShapeSquareBlack = '\uE1B3'
	// Square double whole (4-shape la; Aikin 7-shape la)
	NoteShapeSquareDoubleWhole = '\uECD1'
	// Square white (4-shape la; Aikin 7-shape la)
	NoteShapeSquareWhite = '\uE1B2'
	// Triangle left black (stem up; 4-shape fa; 7-shape fa)
	NoteShapeTriangleLeftBlack = '\uE1B7'
	// Triangle left double whole (stem up; 4-shape fa; 7-shape fa)
	NoteShapeTriangleLeftDoubleWhole = '\uECD3'
	// Triangle left white (stem up; 4-shape fa; 7-shape fa)
	NoteShapeTriangleLeftWhite = '\uE1B6'
	// Triangle right black (stem down; 4-shape fa; 7-shape fa)
	NoteShapeTriangleRightBlack = '\uE1B5'
	// Triangle right double whole (stem down; 4-shape fa; 7-shape fa)
	NoteShapeTriangleRightDoubleWhole = '\uECD2'
	// Triangle right white (stem down; 4-shape fa; 7-shape fa)
	NoteShapeTriangleRightWhite = '\uE1B4'
	// Triangle-round black (Aikin 7-shape ti)
	NoteShapeTriangleRoundBlack = '\uE1BF'
	// Triangle-round white (Aikin 7-shape ti)
	NoteShapeTriangleRoundDoubleWhole = '\uECD7'
	// Triangle-round left black (Funk 7-shape ti)
	NoteShapeTriangleRoundLeftBlack = '\uE1CB'
	// Triangle-round left double whole (Funk 7-shape ti)
	NoteShapeTriangleRoundLeftDoubleWhole = '\uECDD'
	// Triangle-round left white (Funk 7-shape ti)
	NoteShapeTriangleRoundLeftWhite = '\uE1CA'
	// Triangle-round white (Aikin 7-shape ti)
	NoteShapeTriangleRoundWhite = '\uE1BE'
	// Triangle up black (Aikin 7-shape do)
	NoteShapeTriangleUpBlack = '\uE1BB'
	// Triangle up double whole (Aikin 7-shape do)
	NoteShapeTriangleUpDoubleWhole = '\uECD5'
	// Triangle up white (Aikin 7-shape do)
	NoteShapeTriangleUpWhite = '\uE1BA'
	// Si (black note)
	NoteSiBlack = '\uE167'
	// Si (half note)
	NoteSiHalf = '\uE15F'
	// Si (whole note)
	NoteSiWhole = '\uE157'
	// So (black note)
	NoteSoBlack = '\uE164'
	// So (half note)
	NoteSoHalf = '\uE15C'
	// So (whole note)
	NoteSoWhole = '\uE154'
	// Te (black note)
	NoteTeBlack = '\uEEFA'
	// Te (half note)
	NoteTeHalf = '\uEEF1'
	// Te (whole note)
	NoteTeWhole = '\uEEE8'
	// Ti (black note)
	NoteTiBlack = '\uE166'
	// Ti (half note)
	NoteTiHalf = '\uE15E'
	// Ti (whole note)
	NoteTiWhole = '\uE156'
	// Whole note (semibreve)
	NoteWhole = '\uE1D2'
	// Black notehead
	NoteheadBlack = '\uE0A4'
	// Circle slash notehead
	NoteheadCircleSlash = '\uE0F7'
	// Circle X notehead
	NoteheadCircleX = '\uE0B3'
	// Circle X double whole
	NoteheadCircleXDoubleWhole = '\uE0B0'
	// Circle X half
	NoteheadCircleXHalf = '\uE0B2'
	// Circle X whole
	NoteheadCircleXWhole = '\uE0B1'
	// Circled black notehead
	NoteheadCircledBlack = '\uE0E4'
	// Black notehead in large circle
	NoteheadCircledBlackLarge = '\uE0E8'
	// Circled double whole notehead
	NoteheadCircledDoubleWhole = '\uE0E7'
	// Double whole notehead in large circle
	NoteheadCircledDoubleWholeLarge = '\uE0EB'
	// Circled half notehead
	NoteheadCircledHalf = '\uE0E5'
	// Half notehead in large circle
	NoteheadCircledHalfLarge = '\uE0E9'
	// Circled whole notehead
	NoteheadCircledWhole = '\uE0E6'
	// Whole notehead in large circle
	NoteheadCircledWholeLarge = '\uE0EA'
	// Cross notehead in large circle
	NoteheadCircledXLarge = '\uE0EC'
	// Double whole note cluster, 2nd
	NoteheadClusterDoubleWhole2nd = '\uE124'
	// Double whole note cluster, 3rd
	NoteheadClusterDoubleWhole3rd = '\uE128'
	// Combining double whole note cluster, bottom
	NoteheadClusterDoubleWholeBottom = '\uE12E'
	// Combining double whole note cluster, middle
	NoteheadClusterDoubleWholeMiddle = '\uE12D'
	// Combining double whole note cluster, top
	NoteheadClusterDoubleWholeTop = '\uE12C'
	// Half note cluster, 2nd
	NoteheadClusterHalf2nd = '\uE126'
	// Half note cluster, 3rd
	NoteheadClusterHalf3rd = '\uE12A'
	// Combining half note cluster, bottom
	NoteheadClusterHalfBottom = '\uE134'
	// Combining half note cluster, middle
	NoteheadClusterHalfMiddle = '\uE133'
	// Combining half note cluster, top
	NoteheadClusterHalfTop = '\uE132'
	// Quarter note cluster, 2nd
	NoteheadClusterQuarter2nd = '\uE127'
	// Quarter note cluster, 3rd
	NoteheadClusterQuarter3rd = '\uE12B'
	// Combining quarter note cluster, bottom
	NoteheadClusterQuarterBottom = '\uE137'
	// Combining quarter note cluster, middle
	NoteheadClusterQuarterMiddle = '\uE136'
	// Combining quarter note cluster, top
	NoteheadClusterQuarterTop = '\uE135'
	// Cluster notehead black (round)
	NoteheadClusterRoundBlack = '\uE123'
	// Cluster notehead white (round)
	NoteheadClusterRoundWhite = '\uE122'
	// Cluster notehead black (square)
	NoteheadClusterSquareBlack = '\uE121'
	// Cluster notehead white (square)
	NoteheadClusterSquareWhite = '\uE120'
	// Whole note cluster, 2nd
	NoteheadClusterWhole2nd = '\uE125'
	// Whole note cluster, 3rd
	NoteheadClusterWhole3rd = '\uE129'
	// Combining whole note cluster, bottom
	NoteheadClusterWholeBottom = '\uE131'
	// Combining whole note cluster, middle
	NoteheadClusterWholeMiddle = '\uE130'
	// Combining whole note cluster, top
	NoteheadClusterWholeTop = '\uE12F'
	// 4/11 note (eleventh note series, Cowell)
	NoteheadCowellEleventhNoteSeriesHalf = '\uEEAE'
	// 8/11 note (eleventh note series, Cowell)
	NoteheadCowellEleventhNoteSeriesWhole = '\uEEAD'
	// 2/11 note (eleventh note series, Cowell)
	NoteheadCowellEleventhSeriesBlack = '\uEEAF'
	// 2/15 note (fifteenth note series, Cowell)
	NoteheadCowellFifteenthNoteSeriesBlack = '\uEEB5'
	// 4/15 note (fifteenth note series, Cowell)
	NoteheadCowellFifteenthNoteSeriesHalf = '\uEEB4'
	// 8/15 note (fifteenth note series, Cowell)
	NoteheadCowellFifteenthNoteSeriesWhole = '\uEEB3'
	// 1/5 note (fifth note series, Cowell)
	NoteheadCowellFifthNoteSeriesBlack = '\uEEA6'
	// 2/5 note (fifth note series, Cowell)
	NoteheadCowellFifthNoteSeriesHalf = '\uEEA5'
	// 4/5 note (fifth note series, Cowell)
	NoteheadCowellFifthNoteSeriesWhole = '\uEEA4'
	// 2/9 note (ninth note series, Cowell)
	NoteheadCowellNinthNoteSeriesBlack = '\uEEAC'
	// 4/9 note (ninth note series, Cowell)
	NoteheadCowellNinthNoteSeriesHalf = '\uEEAB'
	// 8/9 note (ninth note series, Cowell)
	NoteheadCowellNinthNoteSeriesWhole = '\uEEAA'
	// 1/7 note (seventh note series, Cowell)
	NoteheadCowellSeventhNoteSeriesBlack = '\uEEA9'
	// 2/7 note (seventh note series, Cowell)
	NoteheadCowellSeventhNoteSeriesHalf = '\uEEA8'
	// 4/7 note (seventh note series, Cowell)
	NoteheadCowellSeventhNoteSeriesWhole = '\uEEA7'
	// 1/6 note (third note series, Cowell)
	NoteheadCowellThirdNoteSeriesBlack = '\uEEA3'
	// 1/3 note (third note series, Cowell)
	NoteheadCowellThirdNoteSeriesHalf = '\uEEA2'
	// 2/3 note (third note series, Cowell)
	NoteheadCowellThirdNoteSeriesWhole = '\uEEA1'
	// 2/13 note (thirteenth note series, Cowell)
	NoteheadCowellThirteenthNoteSeriesBlack = '\uEEB2'
	// 4/13 note (thirteenth note series, Cowell)
	NoteheadCowellThirteenthNoteSeriesHalf = '\uEEB1'
	// 8/13 note (thirteenth note series, Cowell)
	NoteheadCowellThirteenthNoteSeriesWhole = '\uEEB0'
	// Diamond black notehead
	NoteheadDiamondBlack = '\uE0DB'
	// Diamond black notehead (old)
	NoteheadDiamondBlackOld = '\uE0E2'
	// Diamond black notehead (wide)
	NoteheadDiamondBlackWide = '\uE0DC'
	// Black diamond cluster, 2nd
	NoteheadDiamondClusterBlack2nd = '\uE139'
	// Black diamond cluster, 3rd
	NoteheadDiamondClusterBlack3rd = '\uE13B'
	// Combining black diamond cluster, bottom
	NoteheadDiamondClusterBlackBottom = '\uE141'
	// Combining black diamond cluster, middle
	NoteheadDiamondClusterBlackMiddle = '\uE140'
	// Combining black diamond cluster, top
	NoteheadDiamondClusterBlackTop = '\uE13F'
	// White diamond cluster, 2nd
	NoteheadDiamondClusterWhite2nd = '\uE138'
	// White diamond cluster, 3rd
	NoteheadDiamondClusterWhite3rd = '\uE13A'
	// Combining white diamond cluster, bottom
	NoteheadDiamondClusterWhiteBottom = '\uE13E'
	// Combining white diamond cluster, middle
	NoteheadDiamondClusterWhiteMiddle = '\uE13D'
	// Combining white diamond cluster, top
	NoteheadDiamondClusterWhiteTop = '\uE13C'
	// Diamond double whole notehead
	NoteheadDiamondDoubleWhole = '\uE0D7'
	// Diamond double whole notehead (old)
	NoteheadDiamondDoubleWholeOld = '\uE0DF'
	// Diamond half notehead
	NoteheadDiamondHalf = '\uE0D9'
	// Half-filled diamond notehead
	NoteheadDiamondHalfFilled = '\uE0E3'
	// Diamond half notehead (old)
	NoteheadDiamondHalfOld = '\uE0E1'
	// Diamond half notehead (wide)
	NoteheadDiamondHalfWide = '\uE0DA'
	// Open diamond notehead
	NoteheadDiamondOpen = '\uE0FC'
	// Diamond white notehead
	NoteheadDiamondWhite = '\uE0DD'
	// Diamond white notehead (wide)
	NoteheadDiamondWhiteWide = '\uE0DE'
	// Diamond whole notehead
	NoteheadDiamondWhole = '\uE0D8'
	// Diamond whole notehead (old)
	NoteheadDiamondWholeOld = '\uE0E0'
	// Double whole (breve) notehead
	NoteheadDoubleWhole = '\uE0A0'
	// Double whole (breve) notehead (square)
	NoteheadDoubleWholeSquare = '\uE0A1'
	// Double whole notehead with X
	NoteheadDoubleWholeWithX = '\uE0B4'
	// Half (minim) notehead
	NoteheadHalf = '\uE0A3'
	// Filled half (minim) notehead
	NoteheadHalfFilled = '\uE0FB'
	// Half notehead with X
	NoteheadHalfWithX = '\uE0B6'
	// Heavy X notehead
	NoteheadHeavyX = '\uE0F8'
	// Heavy X with hat notehead
	NoteheadHeavyXHat = '\uE0F9'
	// Large arrow down (lowest pitch) black notehead
	NoteheadLargeArrowDownBlack = '\uE0F4'
	// Large arrow down (lowest pitch) double whole notehead
	NoteheadLargeArrowDownDoubleWhole = '\uE0F1'
	// Large arrow down (lowest pitch) half notehead
	NoteheadLargeArrowDownHalf = '\uE0F3'
	// Large arrow down (lowest pitch) whole notehead
	NoteheadLargeArrowDownWhole = '\uE0F2'
	// Large arrow up (highest pitch) black notehead
	NoteheadLargeArrowUpBlack = '\uE0F0'
	// Large arrow up (highest pitch) double whole notehead
	NoteheadLargeArrowUpDoubleWhole = '\uE0ED'
	// Large arrow up (highest pitch) half notehead
	NoteheadLargeArrowUpHalf = '\uE0EF'
	// Large arrow up (highest pitch) whole notehead
	NoteheadLargeArrowUpWhole = '\uE0EE'
	// Moon notehead black
	NoteheadMoonBlack = '\uE0CB'
	// Moon notehead white
	NoteheadMoonWhite = '\uE0CA'
	// Sine notehead (Nancarrow)
	NoteheadNancarrowSine = '\uEEA0'
	// Null notehead
	NoteheadNull = '\uE0A5'
	// Parenthesis notehead
	NoteheadParenthesis = '\uE0CE'
	// Opening parenthesis
	NoteheadParenthesisLeft = '\uE0F5'
	// Closing parenthesis
	NoteheadParenthesisRight = '\uE0F6'
	// Plus notehead black
	NoteheadPlusBlack = '\uE0AF'
	// Plus notehead double whole
	NoteheadPlusDoubleWhole = '\uE0AC'
	// Plus notehead half
	NoteheadPlusHalf = '\uE0AE'
	// Plus notehead whole
	NoteheadPlusWhole = '\uE0AD'
	// Combining black rectangular cluster, bottom
	NoteheadRectangularClusterBlackBottom = '\uE144'
	// Combining black rectangular cluster, middle
	NoteheadRectangularClusterBlackMiddle = '\uE143'
	// Combining black rectangular cluster, top
	NoteheadRectangularClusterBlackTop = '\uE142'
	// Combining white rectangular cluster, bottom
	NoteheadRectangularClusterWhiteBottom = '\uE147'
	// Combining white rectangular cluster, middle
	NoteheadRectangularClusterWhiteMiddle = '\uE146'
	// Combining white rectangular cluster, top
	NoteheadRectangularClusterWhiteTop = '\uE145'
	// Round black notehead
	NoteheadRoundBlack = '\uE113'
	// Round black notehead, double slashed
	NoteheadRoundBlackDoubleSlashed = '\uE11C'
	// Large round black notehead
	NoteheadRoundBlackLarge = '\uE110'
	// Round black notehead, slashed
	NoteheadRoundBlackSlashed = '\uE118'
	// Large round black notehead, slashed
	NoteheadRoundBlackSlashedLarge = '\uE116'
	// Round white notehead
	NoteheadRoundWhite = '\uE114'
	// Round white notehead, double slashed
	NoteheadRoundWhiteDoubleSlashed = '\uE11D'
	// Large round white notehead
	NoteheadRoundWhiteLarge = '\uE111'
	// Round white notehead, slashed
	NoteheadRoundWhiteSlashed = '\uE119'
	// Large round white notehead, slashed
	NoteheadRoundWhiteSlashedLarge = '\uE117'
	// Round white notehead with dot
	NoteheadRoundWhiteWithDot = '\uE115'
	// Large round white notehead with dot
	NoteheadRoundWhiteWithDotLarge = '\uE112'
	// Large white diamond
	NoteheadSlashDiamondWhite = '\uE104'
	// Slash with horizontal ends
	NoteheadSlashHorizontalEnds = '\uE101'
	// Muted slash with horizontal ends
	NoteheadSlashHorizontalEndsMuted = '\uE108'
	// Slash with vertical ends
	NoteheadSlashVerticalEnds = '\uE100'
	// Muted slash with vertical ends
	NoteheadSlashVerticalEndsMuted = '\uE107'
	// Small slash with vertical ends
	NoteheadSlashVerticalEndsSmall = '\uE105'
	// White slash double whole
	NoteheadSlashWhiteDoubleWhole = '\uE10A'
	// White slash half
	NoteheadSlashWhiteHalf = '\uE103'
	// Muted white slash
	NoteheadSlashWhiteMuted = '\uE109'
	// White slash whole
	NoteheadSlashWhiteWhole = '\uE102'
	// Large X notehead
	NoteheadSlashX = '\uE106'
	// Slashed black notehead (bottom left to top right)
	NoteheadSlashedBlack1 = '\uE0CF'
	// Slashed black notehead (top left to bottom right)
	NoteheadSlashedBlack2 = '\uE0D0'
	// Slashed double whole notehead (bottom left to top right)
	NoteheadSlashedDoubleWhole1 = '\uE0D5'
	// Slashed double whole notehead (top left to bottom right)
	NoteheadSlashedDoubleWhole2 = '\uE0D6'
	// Slashed half notehead (bottom left to top right)
	NoteheadSlashedHalf1 = '\uE0D1'
	// Slashed half notehead (top left to bottom right)
	NoteheadSlashedHalf2 = '\uE0D2'
	// Slashed whole notehead (bottom left to top right)
	NoteheadSlashedWhole1 = '\uE0D3'
	// Slashed whole notehead (top left to bottom right)
	NoteheadSlashedWhole2 = '\uE0D4'
	// Square notehead black
	NoteheadSquareBlack = '\uE0B9'
	// Large square black notehead
	NoteheadSquareBlackLarge = '\uE11A'
	// Large square white notehead
	NoteheadSquareBlackWhite = '\uE11B'
	// Square notehead white
	NoteheadSquareWhite = '\uE0B8'
	// Triangle notehead down black
	NoteheadTriangleDownBlack = '\uE0C7'
	// Triangle notehead down double whole
	NoteheadTriangleDownDoubleWhole = '\uE0C3'
	// Triangle notehead down half
	NoteheadTriangleDownHalf = '\uE0C5'
	// Triangle notehead down white
	NoteheadTriangleDownWhite = '\uE0C6'
	// Triangle notehead down whole
	NoteheadTriangleDownWhole = '\uE0C4'
	// Triangle notehead left black
	NoteheadTriangleLeftBlack = '\uE0C0'
	// Triangle notehead left white
	NoteheadTriangleLeftWhite = '\uE0BF'
	// Triangle notehead right black
	NoteheadTriangleRightBlack = '\uE0C2'
	// Triangle notehead right white
	NoteheadTriangleRightWhite = '\uE0C1'
	// Triangle-round notehead down black
	NoteheadTriangleRoundDownBlack = '\uE0CD'
	// Triangle-round notehead down white
	NoteheadTriangleRoundDownWhite = '\uE0CC'
	// Triangle notehead up black
	NoteheadTriangleUpBlack = '\uE0BE'
	// Triangle notehead up double whole
	NoteheadTriangleUpDoubleWhole = '\uE0BA'
	// Triangle notehead up half
	NoteheadTriangleUpHalf = '\uE0BC'
	// Triangle notehead up right black
	NoteheadTriangleUpRightBlack = '\uE0C9'
	// Triangle notehead up right white
	NoteheadTriangleUpRightWhite = '\uE0C8'
	// Triangle notehead up white
	NoteheadTriangleUpWhite = '\uE0BD'
	// Triangle notehead up whole
	NoteheadTriangleUpWhole = '\uE0BB'
	// Void notehead with X
	NoteheadVoidWithX = '\uE0B7'
	// Whole (semibreve) notehead
	NoteheadWhole = '\uE0A2'
	// Filled whole (semibreve) notehead
	NoteheadWholeFilled = '\uE0FA'
	// Whole notehead with X
	NoteheadWholeWithX = '\uE0B5'
	// X notehead black
	NoteheadXBlack = '\uE0A9'
	// X notehead double whole
	NoteheadXDoubleWhole = '\uE0A6'
	// X notehead half
	NoteheadXHalf = '\uE0A8'
	// Ornate X notehead
	NoteheadXOrnate = '\uE0AA'
	// Ornate X notehead in ellipse
	NoteheadXOrnateEllipse = '\uE0AB'
	// X notehead whole
	NoteheadXWhole = '\uE0A7'
	// a (baseline)
	OctaveBaselineA = '\uEC91'
	// b (baseline)
	OctaveBaselineB = '\uEC93'
	// m (baseline)
	OctaveBaselineM = '\uEC95'
	// v (baseline)
	OctaveBaselineV = '\uEC97'
	// Bassa
	OctaveBassa = '\uE51F'
	// Loco
	OctaveLoco = '\uEC90'
	// Left parenthesis for octave signs
	OctaveParensLeft = '\uE51A'
	// Right parenthesis for octave signs
	OctaveParensRight = '\uE51B'
	// a (superscript)
	OctaveSuperscriptA = '\uEC92'
	// b (superscript)
	OctaveSuperscriptB = '\uEC94'
	// m (superscript)
	OctaveSuperscriptM = '\uEC96'
	// v (superscript)
	OctaveSuperscriptV = '\uEC98'
	// One-handed roll (Stevens)
	OneHandedRollStevens = '\uE233'
	// Two Fusae
	OrganGerman2Fusae = '\uEE2E'
	// Two Minimae
	OrganGerman2Minimae = '\uEE2C'
	// Combining double octave line above
	OrganGerman2OctaveUp = '\uEE19'
	// Two Semifusae
	OrganGerman2Semifusae = '\uEE2F'
	// Two Semiminimae
	OrganGerman2Semiminimae = '\uEE2D'
	// Three Fusae
	OrganGerman3Fusae = '\uEE32'
	// Three Minimae
	OrganGerman3Minimae = '\uEE30'
	// Three Semifusae
	OrganGerman3Semifusae = '\uEE33'
	// Three Semiminimae
	OrganGerman3Semiminimae = '\uEE31'
	// Four Fusae
	OrganGerman4Fusae = '\uEE36'
	// Four Minimae
	OrganGerman4Minimae = '\uEE34'
	// Four Semifusae
	OrganGerman4Semifusae = '\uEE37'
	// Four Semiminimae
	OrganGerman4Semiminimae = '\uEE35'
	// Five Fusae
	OrganGerman5Fusae = '\uEE3A'
	// Five Minimae
	OrganGerman5Minimae = '\uEE38'
	// Five Semifusae
	OrganGerman5Semifusae = '\uEE3B'
	// Five Semiminimae
	OrganGerman5Semiminimae = '\uEE39'
	// Six Fusae
	OrganGerman6Fusae = '\uEE3E'
	// Six Minimae
	OrganGerman6Minimae = '\uEE3C'
	// Six Semifusae
	OrganGerman6Semifusae = '\uEE3F'
	// Six Semiminimae
	OrganGerman6Semiminimae = '\uEE3D'
	// German organ tablature small A
	OrganGermanALower = '\uEE15'
	// German organ tablature great A
	OrganGermanAUpper = '\uEE09'
	// Rhythm Dot
	OrganGermanAugmentationDot = '\uEE1C'
	// German organ tablature small B
	OrganGermanBLower = '\uEE16'
	// German organ tablature great B
	OrganGermanBUpper = '\uEE0A'
	// Brevis (Binary) Buxheimer Orgelbuch
	OrganGermanBuxheimerBrevis2 = '\uEE25'
	// Brevis (Ternary) Buxheimer Orgelbuch
	OrganGermanBuxheimerBrevis3 = '\uEE24'
	// Minima Rest Buxheimer Orgelbuch
	OrganGermanBuxheimerMinimaRest = '\uEE1E'
	// Semibrevis Buxheimer Orgelbuch
	OrganGermanBuxheimerSemibrevis = '\uEE26'
	// Semibrevis Rest Buxheimer Orgelbuch
	OrganGermanBuxheimerSemibrevisRest = '\uEE1D'
	// German organ tablature small C
	OrganGermanCLower = '\uEE0C'
	// German organ tablature great C
	OrganGermanCUpper = '\uEE00'
	// German organ tablature small Cis
	OrganGermanCisLower = '\uEE0D'
	// German organ tablature great Cis
	OrganGermanCisUpper = '\uEE01'
	// German organ tablature small D
	OrganGermanDLower = '\uEE0E'
	// German organ tablature great D
	OrganGermanDUpper = '\uEE02'
	// German organ tablature small Dis
	OrganGermanDisLower = '\uEE0F'
	// German organ tablature great Dis
	OrganGermanDisUpper = '\uEE03'
	// German organ tablature small E
	OrganGermanELower = '\uEE10'
	// German organ tablature great E
	OrganGermanEUpper = '\uEE04'
	// German organ tablature small F
	OrganGermanFLower = '\uEE11'
	// German organ tablature great F
	OrganGermanFUpper = '\uEE05'
	// German organ tablature small Fis
	OrganGermanFisLower = '\uEE12'
	// German organ tablature great Fis
	OrganGermanFisUpper = '\uEE06'
	// Fusa
	OrganGermanFusa = '\uEE2A'
	// Fusa Rest
	OrganGermanFusaRest = '\uEE22'
	// German organ tablature small G
	OrganGermanGLower = '\uEE13'
	// German organ tablature great G
	OrganGermanGUpper = '\uEE07'
	// German organ tablature small Gis
	OrganGermanGisLower = '\uEE14'
	// German organ tablature great Gis
	OrganGermanGisUpper = '\uEE08'
	// German organ tablature small H
	OrganGermanHLower = '\uEE17'
	// German organ tablature great H
	OrganGermanHUpper = '\uEE0B'
	// Minima
	OrganGermanMinima = '\uEE28'
	// Minima Rest
	OrganGermanMinimaRest = '\uEE20'
	// Combining single octave line below
	OrganGermanOctaveDown = '\uEE1A'
	// Combining single octave line above
	OrganGermanOctaveUp = '\uEE18'
	// Semibrevis
	OrganGermanSemibrevis = '\uEE27'
	// Semibrevis Rest
	OrganGermanSemibrevisRest = '\uEE1F'
	// Semifusa
	OrganGermanSemifusa = '\uEE2B'
	// Semifusa Rest
	OrganGermanSemifusaRest = '\uEE23'
	// Semiminima
	OrganGermanSemiminima = '\uEE29'
	// Semiminima Rest
	OrganGermanSemiminimaRest = '\uEE21'
	// Tie
	OrganGermanTie = '\uEE1B'
	// Ornament bottom left concave stroke
	OrnamentBottomLeftConcaveStroke = '\uE59A'
	// Ornament bottom left concave stroke, large
	OrnamentBottomLeftConcaveStrokeLarge = '\uE59B'
	// Ornament bottom left convex stroke
	OrnamentBottomLeftConvexStroke = '\uE59C'
	// Ornament bottom right concave stroke
	OrnamentBottomRightConcaveStroke = '\uE5A7'
	// Ornament bottom right convex stroke
	OrnamentBottomRightConvexStroke = '\uE5A8'
	// Comma
	OrnamentComma = '\uE581'
	// Double oblique straight lines NW-SE
	OrnamentDoubleObliqueLinesAfterNote = '\uE57E'
	// Double oblique straight lines SW-NE
	OrnamentDoubleObliqueLinesBeforeNote = '\uE57D'
	// Curve below
	OrnamentDownCurve = '\uE578'
	// Haydn ornament
	OrnamentHaydn = '\uE56F'
	// Ornament high left concave stroke
	OrnamentHighLeftConcaveStroke = '\uE592'
	// Ornament high left convex stroke
	OrnamentHighLeftConvexStroke = '\uE593'
	// Ornament high right concave stroke
	OrnamentHighRightConcaveStroke = '\uE5A2'
	// Ornament high right convex stroke
	OrnamentHighRightConvexStroke = '\uE5A3'
	// Hook after note
	OrnamentHookAfterNote = '\uE576'
	// Hook before note
	OrnamentHookBeforeNote = '\uE575'
	// Left-facing half circle
	OrnamentLeftFacingHalfCircle = '\uE572'
	// Left-facing hook
	OrnamentLeftFacingHook = '\uE574'
	// Ornament left +
	OrnamentLeftPlus = '\uE597'
	// Ornament left shake t
	OrnamentLeftShakeT = '\uE596'
	// Ornament left vertical stroke
	OrnamentLeftVerticalStroke = '\uE594'
	// Ornament left vertical stroke with cross (+)
	OrnamentLeftVerticalStrokeWithCross = '\uE595'
	// Ornament low left concave stroke
	OrnamentLowLeftConcaveStroke = '\uE598'
	// Ornament low left convex stroke
	OrnamentLowLeftConvexStroke = '\uE599'
	// Ornament low right concave stroke
	OrnamentLowRightConcaveStroke = '\uE5A5'
	// Ornament low right convex stroke
	OrnamentLowRightConvexStroke = '\uE5A6'
	// Ornament middle vertical stroke
	OrnamentMiddleVerticalStroke = '\uE59F'
	// Mordent
	OrnamentMordent = '\uE56D'
	// Oblique straight line NW-SE
	OrnamentObliqueLineAfterNote = '\uE57C'
	// Oblique straight line SW-NE
	OrnamentObliqueLineBeforeNote = '\uE57B'
	// Oblique straight line tilted NW-SE
	OrnamentObliqueLineHorizAfterNote = '\uE580'
	// Oblique straight line tilted SW-NE
	OrnamentObliqueLineHorizBeforeNote = '\uE57F'
	// Oriscus
	OrnamentOriscus = '\uEA21'
	// Pincé (Couperin)
	OrnamentPinceCouperin = '\uE588'
	// Port de voix
	OrnamentPortDeVoixV = '\uE570'
	// Supported appoggiatura trill
	OrnamentPrecompAppoggTrill = '\uE5B2'
	// Supported appoggiatura trill with two-note suffix
	OrnamentPrecompAppoggTrillSuffix = '\uE5B3'
	// Cadence
	OrnamentPrecompCadence = '\uE5BE'
	// Cadence with upper prefix
	OrnamentPrecompCadenceUpperPrefix = '\uE5C1'
	// Cadence with upper prefix and turn
	OrnamentPrecompCadenceUpperPrefixTurn = '\uE5C2'
	// Cadence with turn
	OrnamentPrecompCadenceWithTurn = '\uE5BF'
	// Descending slide
	OrnamentPrecompDescendingSlide = '\uE5B1'
	// Double cadence with lower prefix
	OrnamentPrecompDoubleCadenceLowerPrefix = '\uE5C0'
	// Double cadence with upper prefix
	OrnamentPrecompDoubleCadenceUpperPrefix = '\uE5C3'
	// Double cadence with upper prefix and turn
	OrnamentPrecompDoubleCadenceUpperPrefixTurn = '\uE5C4'
	// Inverted mordent with upper prefix
	OrnamentPrecompInvertedMordentUpperPrefix = '\uE5C7'
	// Mordent with release
	OrnamentPrecompMordentRelease = '\uE5C5'
	// Mordent with upper prefix
	OrnamentPrecompMordentUpperPrefix = '\uE5C6'
	// Pre-beat port de voix followed by multiple mordent (Dandrieu)
	OrnamentPrecompPortDeVoixMordent = '\uE5BC'
	// Slide
	OrnamentPrecompSlide = '\uE5B0'
	// Slide-trill with two-note suffix (J.S. Bach)
	OrnamentPrecompSlideTrillBach = '\uE5B8'
	// Slide-trill (D'Anglebert)
	OrnamentPrecompSlideTrillDAnglebert = '\uE5B5'
	// Slide-trill with one-note suffix (Marpurg)
	OrnamentPrecompSlideTrillMarpurg = '\uE5B6'
	// Slide-trill (Muffat)
	OrnamentPrecompSlideTrillMuffat = '\uE5B9'
	// Slide-trill with two-note suffix (Muffat)
	OrnamentPrecompSlideTrillSuffixMuffat = '\uE5BA'
	// Trill with lower suffix
	OrnamentPrecompTrillLowerSuffix = '\uE5C8'
	// Trill with two-note suffix (Dandrieu)
	OrnamentPrecompTrillSuffixDandrieu = '\uE5BB'
	// Trill with mordent
	OrnamentPrecompTrillWithMordent = '\uE5BD'
	// Turn-trill with two-note suffix (J.S. Bach)
	OrnamentPrecompTurnTrillBach = '\uE5B7'
	// Turn-trill (D'Anglebert)
	OrnamentPrecompTurnTrillDAnglebert = '\uE5B4'
	// Quilisma
	OrnamentQuilisma = '\uEA20'
	// Right-facing half circle
	OrnamentRightFacingHalfCircle = '\uE571'
	// Right-facing hook
	OrnamentRightFacingHook = '\uE573'
	// Ornament right vertical stroke
	OrnamentRightVerticalStroke = '\uE5A4'
	// Schleifer (long mordent)
	OrnamentSchleifer = '\uE587'
	// Shake
	OrnamentShake3 = '\uE582'
	// Shake (Muffat)
	OrnamentShakeMuffat1 = '\uE584'
	// Short oblique straight line NW-SE
	OrnamentShortObliqueLineAfterNote = '\uE57A'
	// Short oblique straight line SW-NE
	OrnamentShortObliqueLineBeforeNote = '\uE579'
	// Short trill
	OrnamentShortTrill = '\uE56C'
	// Ornament top left concave stroke
	OrnamentTopLeftConcaveStroke = '\uE590'
	// Ornament top left convex stroke
	OrnamentTopLeftConvexStroke = '\uE591'
	// Ornament top right concave stroke
	OrnamentTopRightConcaveStroke = '\uE5A0'
	// Ornament top right convex stroke
	OrnamentTopRightConvexStroke = '\uE5A1'
	// Tremblement
	OrnamentTremblement = '\uE56E'
	// Tremblement appuyé (Couperin)
	OrnamentTremblementCouperin = '\uE589'
	// Trill
	OrnamentTrill = '\uE566'
	// Turn
	OrnamentTurn = '\uE567'
	// Inverted turn
	OrnamentTurnInverted = '\uE568'
	// Turn with slash
	OrnamentTurnSlash = '\uE569'
	// Turn up
	OrnamentTurnUp = '\uE56A'
	// Inverted turn up
	OrnamentTurnUpS = '\uE56B'
	// Curve above
	OrnamentUpCurve = '\uE577'
	// Vertical line
	OrnamentVerticalLine = '\uE583'
	// Ornament zig-zag line without right-hand end
	OrnamentZigZagLineNoRightEnd = '\uE59D'
	// Ornament zig-zag line with right-hand end
	OrnamentZigZagLineWithRightEnd = '\uE59E'
	// Ottava
	Ottava = '\uE510'
	// Ottava alta
	OttavaAlta = '\uE511'
	// Ottava bassa
	OttavaBassa = '\uE512'
	// Ottava bassa (ba)
	OttavaBassaBa = '\uE513'
	// Ottava bassa (8vb)
	OttavaBassaVb = '\uE51C'
	// Penderecki unmeasured tremolo
	PendereckiTremolo = '\uE22B'
	// Agogo
	PictAgogo = '\uE717'
	// Almglocken
	PictAlmglocken = '\uE712'
	// Anvil
	PictAnvil = '\uE701'
	// Bamboo tube chimes
	PictBambooChimes = '\uE6C3'
	// Bamboo scraper
	PictBambooScraper = '\uE6FB'
	// Bass drum
	PictBassDrum = '\uE6D4'
	// Bass drum on side
	PictBassDrumOnSide = '\uE6D5'
	// Bow
	PictBeaterBow = '\uE7DE'
	// Box for percussion beater
	PictBeaterBox = '\uE7EB'
	// Brass mallets down
	PictBeaterBrassMalletsDown = '\uE7DA'
	// Brass mallets left
	PictBeaterBrassMalletsLeft = '\uE7EE'
	// Brass mallets right
	PictBeaterBrassMalletsRight = '\uE7ED'
	// Brass mallets up
	PictBeaterBrassMalletsUp = '\uE7D9'
	// Combining dashed circle for round beaters (plated)
	PictBeaterCombiningDashedCircle = '\uE7EA'
	// Combining parentheses for round beaters (padded)
	PictBeaterCombiningParentheses = '\uE7E9'
	// Double bass drum stick down
	PictBeaterDoubleBassDrumDown = '\uE7A1'
	// Double bass drum stick up
	PictBeaterDoubleBassDrumUp = '\uE7A0'
	// Finger
	PictBeaterFinger = '\uE7E4'
	// Fingernails
	PictBeaterFingernails = '\uE7E6'
	// Fist
	PictBeaterFist = '\uE7E5'
	// Guiro scraper
	PictBeaterGuiroScraper = '\uE7DD'
	// Hammer
	PictBeaterHammer = '\uE7E1'
	// Metal hammer, down
	PictBeaterHammerMetalDown = '\uE7D0'
	// Metal hammer, up
	PictBeaterHammerMetalUp = '\uE7CF'
	// Plastic hammer, down
	PictBeaterHammerPlasticDown = '\uE7CE'
	// Plastic hammer, up
	PictBeaterHammerPlasticUp = '\uE7CD'
	// Wooden hammer, down
	PictBeaterHammerWoodDown = '\uE7CC'
	// Wooden hammer, up
	PictBeaterHammerWoodUp = '\uE7CB'
	// Hand
	PictBeaterHand = '\uE7E3'
	// Hard bass drum stick down
	PictBeaterHardBassDrumDown = '\uE79D'
	// Hard bass drum stick up
	PictBeaterHardBassDrumUp = '\uE79C'
	// Hard glockenspiel stick down
	PictBeaterHardGlockenspielDown = '\uE785'
	// Hard glockenspiel stick left
	PictBeaterHardGlockenspielLeft = '\uE787'
	// Hard glockenspiel stick right
	PictBeaterHardGlockenspielRight = '\uE786'
	// Hard glockenspiel stick up
	PictBeaterHardGlockenspielUp = '\uE784'
	// Hard timpani stick down
	PictBeaterHardTimpaniDown = '\uE791'
	// Hard timpani stick left
	PictBeaterHardTimpaniLeft = '\uE793'
	// Hard timpani stick right
	PictBeaterHardTimpaniRight = '\uE792'
	// Hard timpani stick up
	PictBeaterHardTimpaniUp = '\uE790'
	// Hard xylophone stick down
	PictBeaterHardXylophoneDown = '\uE779'
	// Hard xylophone stick left
	PictBeaterHardXylophoneLeft = '\uE77B'
	// Hard xylophone stick right
	PictBeaterHardXylophoneRight = '\uE77A'
	// Hard xylophone stick up
	PictBeaterHardXylophoneUp = '\uE778'
	// Hard yarn beater down
	PictBeaterHardYarnDown = '\uE7AB'
	// Hard yarn beater left
	PictBeaterHardYarnLeft = '\uE7AD'
	// Hard yarn beater right
	PictBeaterHardYarnRight = '\uE7AC'
	// Hard yarn beater up
	PictBeaterHardYarnUp = '\uE7AA'
	// Jazz sticks down
	PictBeaterJazzSticksDown = '\uE7D4'
	// Jazz sticks up
	PictBeaterJazzSticksUp = '\uE7D3'
	// Knitting needle
	PictBeaterKnittingNeedle = '\uE7E2'
	// Chime hammer up
	PictBeaterMallet = '\uE7DF'
	// Chime hammer down
	PictBeaterMalletDown = '\uE7EC'
	// Medium bass drum stick down
	PictBeaterMediumBassDrumDown = '\uE79B'
	// Medium bass drum stick up
	PictBeaterMediumBassDrumUp = '\uE79A'
	// Medium timpani stick down
	PictBeaterMediumTimpaniDown = '\uE78D'
	// Medium timpani stick left
	PictBeaterMediumTimpaniLeft = '\uE78F'
	// Medium timpani stick right
	PictBeaterMediumTimpaniRight = '\uE78E'
	// Medium timpani stick up
	PictBeaterMediumTimpaniUp = '\uE78C'
	// Medium xylophone stick down
	PictBeaterMediumXylophoneDown = '\uE775'
	// Medium xylophone stick left
	PictBeaterMediumXylophoneLeft = '\uE777'
	// Medium xylophone stick right
	PictBeaterMediumXylophoneRight = '\uE776'
	// Medium xylophone stick up
	PictBeaterMediumXylophoneUp = '\uE774'
	// Medium yarn beater down
	PictBeaterMediumYarnDown = '\uE7A7'
	// Medium yarn beater left
	PictBeaterMediumYarnLeft = '\uE7A9'
	// Medium yarn beater right
	PictBeaterMediumYarnRight = '\uE7A8'
	// Medium yarn beater up
	PictBeaterMediumYarnUp = '\uE7A6'
	// Metal bass drum stick down
	PictBeaterMetalBassDrumDown = '\uE79F'
	// Metal bass drum stick up
	PictBeaterMetalBassDrumUp = '\uE79E'
	// Metal beater down
	PictBeaterMetalDown = '\uE7C8'
	// Metal hammer
	PictBeaterMetalHammer = '\uE7E0'
	// Metal beater, left
	PictBeaterMetalLeft = '\uE7CA'
	// Metal beater, right
	PictBeaterMetalRight = '\uE7C9'
	// Metal beater, up
	PictBeaterMetalUp = '\uE7C7'
	// Snare sticks down
	PictBeaterSnareSticksDown = '\uE7D2'
	// Snare sticks up
	PictBeaterSnareSticksUp = '\uE7D1'
	// Soft bass drum stick down
	PictBeaterSoftBassDrumDown = '\uE799'
	// Soft bass drum stick up
	PictBeaterSoftBassDrumUp = '\uE798'
	// Soft glockenspiel stick down
	PictBeaterSoftGlockenspielDown = '\uE781'
	// Soft glockenspiel stick left
	PictBeaterSoftGlockenspielLeft = '\uE783'
	// Soft glockenspiel stick right
	PictBeaterSoftGlockenspielRight = '\uE782'
	// Soft glockenspiel stick up
	PictBeaterSoftGlockenspielUp = '\uE780'
	// Soft timpani stick down
	PictBeaterSoftTimpaniDown = '\uE789'
	// Soft timpani stick left
	PictBeaterSoftTimpaniLeft = '\uE78B'
	// Soft timpani stick right
	PictBeaterSoftTimpaniRight = '\uE78A'
	// Soft timpani stick up
	PictBeaterSoftTimpaniUp = '\uE788'
	// Soft xylophone beaters
	PictBeaterSoftXylophone = '\uE7DB'
	// Soft xylophone stick down
	PictBeaterSoftXylophoneDown = '\uE771'
	// Soft xylophone stick left
	PictBeaterSoftXylophoneLeft = '\uE773'
	// Soft xylophone stick right
	PictBeaterSoftXylophoneRight = '\uE772'
	// Soft xylophone stick up
	PictBeaterSoftXylophoneUp = '\uE770'
	// Soft yarn beater down
	PictBeaterSoftYarnDown = '\uE7A3'
	// Soft yarn beater left
	PictBeaterSoftYarnLeft = '\uE7A5'
	// Soft yarn beater right
	PictBeaterSoftYarnRight = '\uE7A4'
	// Soft yarn beater up
	PictBeaterSoftYarnUp = '\uE7A2'
	// Spoon-shaped wooden mallet
	PictBeaterSpoonWoodenMallet = '\uE7DC'
	// Superball beater down
	PictBeaterSuperballDown = '\uE7AF'
	// Superball beater left
	PictBeaterSuperballLeft = '\uE7B1'
	// Superball beater right
	PictBeaterSuperballRight = '\uE7B0'
	// Superball beater up
	PictBeaterSuperballUp = '\uE7AE'
	// Triangle beater down
	PictBeaterTriangleDown = '\uE7D6'
	// Triangle beater plain
	PictBeaterTrianglePlain = '\uE7EF'
	// Triangle beater up
	PictBeaterTriangleUp = '\uE7D5'
	// Wire brushes down
	PictBeaterWireBrushesDown = '\uE7D8'
	// Wire brushes up
	PictBeaterWireBrushesUp = '\uE7D7'
	// Wood timpani stick down
	PictBeaterWoodTimpaniDown = '\uE795'
	// Wood timpani stick left
	PictBeaterWoodTimpaniLeft = '\uE797'
	// Wood timpani stick right
	PictBeaterWoodTimpaniRight = '\uE796'
	// Wood timpani stick up
	PictBeaterWoodTimpaniUp = '\uE794'
	// Wood xylophone stick down
	PictBeaterWoodXylophoneDown = '\uE77D'
	// Wood xylophone stick left
	PictBeaterWoodXylophoneLeft = '\uE77F'
	// Wood xylophone stick right
	PictBeaterWoodXylophoneRight = '\uE77E'
	// Wood xylophone stick up
	PictBeaterWoodXylophoneUp = '\uE77C'
	// Bell
	PictBell = '\uE714'
	// Bell of cymbal
	PictBellOfCymbal = '\uE72A'
	// Bell plate
	PictBellPlate = '\uE713'
	// Bell tree
	PictBellTree = '\uE71A'
	// Bird whistle
	PictBirdWhistle = '\uE751'
	// Board clapper
	PictBoardClapper = '\uE6F7'
	// Bongos
	PictBongos = '\uE6DD'
	// Brake drum
	PictBrakeDrum = '\uE6E1'
	// Cabasa
	PictCabasa = '\uE743'
	// Cannon
	PictCannon = '\uE761'
	// Car horn
	PictCarHorn = '\uE755'
	// Castanets
	PictCastanets = '\uE6F8'
	// Castanets with handle
	PictCastanetsWithHandle = '\uE6F9'
	// Celesta
	PictCelesta = '\uE6B0'
	// Cencerro
	PictCencerro = '\uE716'
	// Center (Weinberg)
	PictCenter1 = '\uE7FE'
	// Center (Ghent)
	PictCenter2 = '\uE7FF'
	// Center (Caltabiano)
	PictCenter3 = '\uE800'
	// Chain rattle
	PictChainRattle = '\uE748'
	// Chimes
	PictChimes = '\uE6C2'
	// Chinese cymbal
	PictChineseCymbal = '\uE726'
	// Choke (Weinberg)
	PictChokeCymbal = '\uE805'
	// Claves
	PictClaves = '\uE6F2'
	// Coins
	PictCoins = '\uE7E7'
	// Conga
	PictConga = '\uE6DE'
	// Cow bell
	PictCowBell = '\uE711'
	// Crash cymbals
	PictCrashCymbals = '\uE720'
	// Crotales
	PictCrotales = '\uE6AE'
	// Combining crush for stem
	PictCrushStem = '\uE80C'
	// Cuica
	PictCuica = '\uE6E4'
	// Cymbal tongs
	PictCymbalTongs = '\uE728'
	// Damp
	PictDamp1 = '\uE7F9'
	// Damp 2
	PictDamp2 = '\uE7FA'
	// Damp 3
	PictDamp3 = '\uE7FB'
	// Damp 4
	PictDamp4 = '\uE7FC'
	// Combining X for stem (dead note)
	PictDeadNoteStem = '\uE80D'
	// Drum stick
	PictDrumStick = '\uE7E8'
	// Duck call
	PictDuckCall = '\uE757'
	// Edge of cymbal
	PictEdgeOfCymbal = '\uE729'
	// Empty trapezoid
	PictEmptyTrap = '\uE6A9'
	// Finger cymbals
	PictFingerCymbals = '\uE727'
	// Flexatone
	PictFlexatone = '\uE740'
	// Football rattle
	PictFootballRatchet = '\uE6F5'
	// Glass harmonica
	PictGlassHarmonica = '\uE765'
	// Glass harp
	PictGlassHarp = '\uE764'
	// Glass plate chimes
	PictGlassPlateChimes = '\uE6C6'
	// Glass tube chimes
	PictGlassTubeChimes = '\uE6C5'
	// Glockenspiel
	PictGlsp = '\uE6A0'
	// Glockenspiel (Smith Brindle)
	PictGlspSmithBrindle = '\uE6AA'
	// Goblet drum (djembe, dumbek)
	PictGobletDrum = '\uE6E2'
	// Gong
	PictGong = '\uE732'
	// Gong with button (nipple)
	PictGongWithButton = '\uE733'
	// Guiro
	PictGuiro = '\uE6F3'
	// Hard gum beater, down
	PictGumHardDown = '\uE7C4'
	// Hard gum beater, left
	PictGumHardLeft = '\uE7C6'
	// Hard gum beater, right
	PictGumHardRight = '\uE7C5'
	// Hard gum beater, up
	PictGumHardUp = '\uE7C3'
	// Medium gum beater, down
	PictGumMediumDown = '\uE7C0'
	// Medium gum beater, left
	PictGumMediumLeft = '\uE7C2'
	// Medium gum beater, right
	PictGumMediumRight = '\uE7C1'
	// Medium gum beater, up
	PictGumMediumUp = '\uE7BF'
	// Soft gum beater, down
	PictGumSoftDown = '\uE7BC'
	// Soft gum beater, left
	PictGumSoftLeft = '\uE7BE'
	// Soft gum beater, right
	PictGumSoftRight = '\uE7BD'
	// Soft gum beater, up
	PictGumSoftUp = '\uE7BB'
	// Half-open
	PictHalfOpen1 = '\uE7F6'
	// Half-open 2 (Weinberg)
	PictHalfOpen2 = '\uE7F7'
	// Handbell
	PictHandbell = '\uE715'
	// Hi-hat
	PictHiHat = '\uE722'
	// Hi-hat cymbals on stand
	PictHiHatOnStand = '\uE723'
	// Jaw harp
	PictJawHarp = '\uE767'
	// Jingle bells
	PictJingleBells = '\uE719'
	// Klaxon horn
	PictKlaxonHorn = '\uE756'
	// Right hand (Agostini)
	PictLeftHandCircle = '\uE807'
	// Lion's roar
	PictLionsRoar = '\uE763'
	// Lithophone
	PictLithophone = '\uE6B1'
	// Log drum
	PictLogDrum = '\uE6DF'
	// Lotus flute
	PictLotusFlute = '\uE75A'
	// Marimba
	PictMar = '\uE6A6'
	// Marimba (Smith Brindle)
	PictMarSmithBrindle = '\uE6AC'
	// Maraca
	PictMaraca = '\uE741'
	// Maracas
	PictMaracas = '\uE742'
	// Megaphone
	PictMegaphone = '\uE759'
	// Metal plate chimes
	PictMetalPlateChimes = '\uE6C8'
	// Metal tube chimes
	PictMetalTubeChimes = '\uE6C7'
	// Musical saw
	PictMusicalSaw = '\uE766'
	// Normal position (Caltabiano)
	PictNormalPosition = '\uE804'
	// On rim
	PictOnRim = '\uE7F4'
	// Open
	PictOpen = '\uE7F8'
	// Closed / rim shot
	PictOpenRimShot = '\uE7F5'
	// Pistol shot
	PictPistolShot = '\uE760'
	// Police whistle
	PictPoliceWhistle = '\uE752'
	// Quijada (jawbone)
	PictQuijada = '\uE6FA'
	// Rainstick
	PictRainstick = '\uE747'
	// Ratchet
	PictRatchet = '\uE6F4'
	// Reco-reco
	PictRecoReco = '\uE6FC'
	// Left hand (Agostini)
	PictRightHandSquare = '\uE806'
	// Rim or edge (Weinberg)
	PictRim1 = '\uE801'
	// Rim (Ghent)
	PictRim2 = '\uE802'
	// Rim (Caltabiano)
	PictRim3 = '\uE803'
	// Rim shot for stem
	PictRimShotOnStem = '\uE7FD'
	// Sandpaper blocks
	PictSandpaperBlocks = '\uE762'
	// Scrape around rim (counter-clockwise)
	PictScrapeAroundRim = '\uE7F3'
	// Scrape around rim (clockwise)
	PictScrapeAroundRimClockwise = '\uE80E'
	// Scrape from center to edge
	PictScrapeCenterToEdge = '\uE7F1'
	// Scrape from edge to center
	PictScrapeEdgeToCenter = '\uE7F2'
	// Shell bells
	PictShellBells = '\uE718'
	// Shell chimes
	PictShellChimes = '\uE6C4'
	// Siren
	PictSiren = '\uE753'
	// Sistrum
	PictSistrum = '\uE746'
	// Sizzle cymbal
	PictSizzleCymbal = '\uE724'
	// Sleigh bell
	PictSleighBell = '\uE710'
	// Slide brush on gong
	PictSlideBrushOnGong = '\uE734'
	// Slide whistle
	PictSlideWhistle = '\uE750'
	// Slit drum
	PictSlitDrum = '\uE6E0'
	// Snare drum
	PictSnareDrum = '\uE6D1'
	// Military snare drum
	PictSnareDrumMilitary = '\uE6D3'
	// Snare drum, snares off
	PictSnareDrumSnaresOff = '\uE6D2'
	// Steel drums
	PictSteelDrums = '\uE6AF'
	// Stick shot
	PictStickShot = '\uE7F0'
	// Superball
	PictSuperball = '\uE7B2'
	// Suspended cymbal
	PictSuspendedCymbal = '\uE721'
	// Combining swish for stem
	PictSwishStem = '\uE808'
	// Indian tabla
	PictTabla = '\uE6E3'
	// Tam-tam
	PictTamTam = '\uE730'
	// Tam-tam with beater (Smith Brindle)
	PictTamTamWithBeater = '\uE731'
	// Tambourine
	PictTambourine = '\uE6DB'
	// Temple blocks
	PictTempleBlocks = '\uE6F1'
	// Tenor drum
	PictTenorDrum = '\uE6D6'
	// Thundersheet
	PictThundersheet = '\uE744'
	// Timbales
	PictTimbales = '\uE6DC'
	// Timpani
	PictTimpani = '\uE6D0'
	// Tom-tom
	PictTomTom = '\uE6D7'
	// Chinese tom-tom
	PictTomTomChinese = '\uE6D8'
	// Indo-American tom tom
	PictTomTomIndoAmerican = '\uE6DA'
	// Japanese tom-tom
	PictTomTomJapanese = '\uE6D9'
	// Triangle
	PictTriangle = '\uE700'
	// Tubaphone
	PictTubaphone = '\uE6B2'
	// Tubular bells
	PictTubularBells = '\uE6C0'
	// Combining turn left for stem
	PictTurnLeftStem = '\uE80A'
	// Combining turn left or right for stem
	PictTurnRightLeftStem = '\uE80B'
	// Combining turn right for stem
	PictTurnRightStem = '\uE809'
	// Vibraphone
	PictVib = '\uE6A7'
	// Metallophone (vibraphone motor off)
	PictVibMotorOff = '\uE6A8'
	// Vibraphone (Smith Brindle)
	PictVibSmithBrindle = '\uE6AD'
	// Vibraslap
	PictVibraslap = '\uE745'
	// Vietnamese hat cymbal
	PictVietnameseHat = '\uE725'
	// Whip
	PictWhip = '\uE6F6'
	// Wind chimes (glass)
	PictWindChimesGlass = '\uE6C1'
	// Wind machine
	PictWindMachine = '\uE754'
	// Wind whistle (or mouth siren)
	PictWindWhistle = '\uE758'
	// Wood block
	PictWoodBlock = '\uE6F0'
	// Wound beater, hard core down
	PictWoundHardDown = '\uE7B4'
	// Wound beater, hard core left
	PictWoundHardLeft = '\uE7B6'
	// Wound beater, hard core right
	PictWoundHardRight = '\uE7B5'
	// Wound beater, hard core up
	PictWoundHardUp = '\uE7B3'
	// Wound beater, soft core down
	PictWoundSoftDown = '\uE7B8'
	// Wound beater, soft core left
	PictWoundSoftLeft = '\uE7BA'
	// Wound beater, soft core right
	PictWoundSoftRight = '\uE7B9'
	// Wound beater, soft core up
	PictWoundSoftUp = '\uE7B7'
	// Xylophone
	PictXyl = '\uE6A1'
	// Bass xylophone
	PictXylBass = '\uE6A3'
	// Xylophone (Smith Brindle)
	PictXylSmithBrindle = '\uE6AB'
	// Tenor xylophone
	PictXylTenor = '\uE6A2'
	// Trough tenor xylophone
	PictXylTenorTrough = '\uE6A5'
	// Trough xylophone
	PictXylTrough = '\uE6A4'
	// Buzz pizzicato
	PluckedBuzzPizzicato = '\uE632'
	// Damp
	PluckedDamp = '\uE638'
	// Damp all
	PluckedDampAll = '\uE639'
	// Damp for stem
	PluckedDampOnStem = '\uE63B'
	// Fingernail flick
	PluckedFingernailFlick = '\uE637'
	// Left-hand pizzicato
	PluckedLeftHandPizzicato = '\uE633'
	// Plectrum
	PluckedPlectrum = '\uE63A'
	// Snap pizzicato above
	PluckedSnapPizzicatoAbove = '\uE631'
	// Snap pizzicato below
	PluckedSnapPizzicatoBelow = '\uE630'
	// With fingernails
	PluckedWithFingernails = '\uE636'
	// Quindicesima
	Quindicesima = '\uE514'
	// Quindicesima alta
	QuindicesimaAlta = '\uE515'
	// Quindicesima bassa
	QuindicesimaBassa = '\uE516'
	// Quindicesima bassa (mb)
	QuindicesimaBassaMb = '\uE51D'
	// Repeat last bar
	Repeat1Bar = '\uE500'
	// Repeat last two bars
	Repeat2Bars = '\uE501'
	// Repeat last four bars
	Repeat4Bars = '\uE502'
	// Repeat bar lower dot
	RepeatBarLowerDot = '\uE505'
	// Repeat bar slash
	RepeatBarSlash = '\uE504'
	// Repeat bar upper dot
	RepeatBarUpperDot = '\uE503'
	// Single repeat dot
	RepeatDot = '\uE044'
	// Repeat dots
	RepeatDots = '\uE043'
	// Left (start) repeat sign
	RepeatLeft = '\uE040'
	// Right (end) repeat sign
	RepeatRight = '\uE041'
	// Right and left repeat sign
	RepeatRightLeft = '\uE042'
	// 1024th rest
	Rest1024th = '\uE4ED'
	// 128th (semihemidemisemiquaver) rest
	Rest128th = '\uE4EA'
	// 16th (semiquaver) rest
	Rest16th = '\uE4E7'
	// 256th rest
	Rest256th = '\uE4EB'
	// 32nd (demisemiquaver) rest
	Rest32nd = '\uE4E8'
	// 512th rest
	Rest512th = '\uE4EC'
	// 64th (hemidemisemiquaver) rest
	Rest64th = '\uE4E9'
	// Eighth (quaver) rest
	Rest8th = '\uE4E6'
	// Double whole (breve) rest
	RestDoubleWhole = '\uE4E2'
	// Double whole rest on leger lines
	RestDoubleWholeLegerLine = '\uE4F3'
	// Multiple measure rest
	RestHBar = '\uE4EE'
	// H-bar, left half
	RestHBarLeft = '\uE4EF'
	// H-bar, middle
	RestHBarMiddle = '\uE4F0'
	// H-bar, right half
	RestHBarRight = '\uE4F1'
	// Half (minim) rest
	RestHalf = '\uE4E4'
	// Half rest on leger line
	RestHalfLegerLine = '\uE4F5'
	// Longa rest
	RestLonga = '\uE4E1'
	// Maxima rest
	RestMaxima = '\uE4E0'
	// Quarter (crotchet) rest
	RestQuarter = '\uE4E5'
	// Old-style quarter (crotchet) rest
	RestQuarterOld = '\uE4F2'
	// Z-style quarter (crotchet) rest
	RestQuarterZ = '\uE4F6'
	// Whole (semibreve) rest
	RestWhole = '\uE4E3'
	// Whole rest on leger line
	RestWholeLegerLine = '\uE4F4'
	// Reversed brace
	ReversedBrace = '\uE001'
	// Reversed bracket bottom
	ReversedBracketBottom = '\uE006'
	// Reversed bracket top
	ReversedBracketTop = '\uE005'
	// Right repeat sign within bar
	RightRepeatSmall = '\uE04D'
	// Scale degree 1
	ScaleDegree1 = '\uEF00'
	// Scale degree 2
	ScaleDegree2 = '\uEF01'
	// Scale degree 3
	ScaleDegree3 = '\uEF02'
	// Scale degree 4
	ScaleDegree4 = '\uEF03'
	// Scale degree 5
	ScaleDegree5 = '\uEF04'
	// Scale degree 6
	ScaleDegree6 = '\uEF05'
	// Scale degree 7
	ScaleDegree7 = '\uEF06'
	// Scale degree 8
	ScaleDegree8 = '\uEF07'
	// Scale degree 9
	ScaleDegree9 = '\uEF08'
	// Schäffer clef
	SchaefferClef = '\uE06F'
	// Schäffer F clef to G clef change
	SchaefferFClefToGClef = '\uE072'
	// Schäffer G clef to F clef change
	SchaefferGClefToFClef = '\uE071'
	// Schäffer previous clef
	SchaefferPreviousClef = '\uE070'
	// Segno
	Segno = '\uE047'
	// Segno (serpent)
	SegnoSerpent1 = '\uE04A'
	// Segno (serpent with vertical lines)
	SegnoSerpent2 = '\uE04B'
	// Semi-pitched percussion clef 1
	SemipitchedPercussionClef1 = '\uE06B'
	// Semi-pitched percussion clef 2
	SemipitchedPercussionClef2 = '\uE06C'
	// Flat
	SmnFlat = '\uEC52'
	// Flat (white)
	SmnFlatWhite = '\uEC53'
	// Double flat history sign
	SmnHistoryDoubleFlat = '\uEC57'
	// Double sharp history sign
	SmnHistoryDoubleSharp = '\uEC55'
	// Flat history sign
	SmnHistoryFlat = '\uEC56'
	// Sharp history sign
	SmnHistorySharp = '\uEC54'
	// Natural (N)
	SmnNatural = '\uEC58'
	// Sharp stem up
	SmnSharp = '\uEC50'
	// Sharp stem down
	SmnSharpDown = '\uEC59'
	// Sharp (white) stem up
	SmnSharpWhite = '\uEC51'
	// Sharp (white) stem down
	SmnSharpWhiteDown = '\uEC5A'
	// Split bar divider (bar spans a system break)
	SplitBarDivider = '\uE00A'
	// 1-line staff
	Staff1Line = '\uE010'
	// 1-line staff (narrow)
	Staff1LineNarrow = '\uE01C'
	// 1-line staff (wide)
	Staff1LineWide = '\uE016'
	// 2-line staff
	Staff2Lines = '\uE011'
	// 2-line staff (narrow)
	Staff2LinesNarrow = '\uE01D'
	// 2-line staff (wide)
	Staff2LinesWide = '\uE017'
	// 3-line staff
	Staff3Lines = '\uE012'
	// 3-line staff (narrow)
	Staff3LinesNarrow = '\uE01E'
	// 3-line staff (wide)
	Staff3LinesWide = '\uE018'
	// 4-line staff
	Staff4Lines = '\uE013'
	// 4-line staff (narrow)
	Staff4LinesNarrow = '\uE01F'
	// 4-line staff (wide)
	Staff4LinesWide = '\uE019'
	// 5-line staff
	Staff5Lines = '\uE014'
	// 5-line staff (narrow)
	Staff5LinesNarrow = '\uE020'
	// 5-line staff (wide)
	Staff5LinesWide = '\uE01A'
	// 6-line staff
	Staff6Lines = '\uE015'
	// 6-line staff (narrow)
	Staff6LinesNarrow = '\uE021'
	// 6-line staff (wide)
	Staff6LinesWide = '\uE01B'
	// Staff divide arrow down
	StaffDivideArrowDown = '\uE00B'
	// Staff divide arrow up
	StaffDivideArrowUp = '\uE00C'
	// Staff divide arrows
	StaffDivideArrowUpDown = '\uE00D'
	// Lower 1 staff position
	StaffPosLower1 = '\uEB98'
	// Lower 2 staff positions
	StaffPosLower2 = '\uEB99'
	// Lower 3 staff positions
	StaffPosLower3 = '\uEB9A'
	// Lower 4 staff positions
	StaffPosLower4 = '\uEB9B'
	// Lower 5 staff positions
	StaffPosLower5 = '\uEB9C'
	// Lower 6 staff positions
	StaffPosLower6 = '\uEB9D'
	// Lower 7 staff positions
	StaffPosLower7 = '\uEB9E'
	// Lower 8 staff positions
	StaffPosLower8 = '\uEB9F'
	// Raise 1 staff position
	StaffPosRaise1 = '\uEB90'
	// Raise 2 staff positions
	StaffPosRaise2 = '\uEB91'
	// Raise 3 staff positions
	StaffPosRaise3 = '\uEB92'
	// Raise 4 staff positions
	StaffPosRaise4 = '\uEB93'
	// Raise 5 staff positions
	StaffPosRaise5 = '\uEB94'
	// Raise 6 staff positions
	StaffPosRaise6 = '\uEB95'
	// Raise 7 staff positions
	StaffPosRaise7 = '\uEB96'
	// Raise 8 staff positions
	StaffPosRaise8 = '\uEB97'
	// Combining stem
	Stem = '\uE210'
	// Combining bow on bridge stem
	StemBowOnBridge = '\uE215'
	// Combining bow on tailpiece stem
	StemBowOnTailpiece = '\uE216'
	// Combining buzz roll stem
	StemBuzzRoll = '\uE217'
	// Combining damp stem
	StemDamp = '\uE218'
	// Combining harp string noise stem
	StemHarpStringNoise = '\uE21F'
	// Combining multiphonics (black) stem
	StemMultiphonicsBlack = '\uE21A'
	// Combining multiphonics (black and white) stem
	StemMultiphonicsBlackWhite = '\uE21C'
	// Combining multiphonics (white) stem
	StemMultiphonicsWhite = '\uE21B'
	// Combining Penderecki unmeasured tremolo stem
	StemPendereckiTremolo = '\uE213'
	// Combining rim shot stem
	StemRimShot = '\uE21E'
	// Combining sprechgesang stem
	StemSprechgesang = '\uE211'
	// Combining sul ponticello (bow behind bridge) stem
	StemSulPonticello = '\uE214'
	// Combining sussurando stem
	StemSussurando = '\uE21D'
	// Combining swished stem
	StemSwished = '\uE212'
	// Combining vibrato pulse accent (Saunders) stem
	StemVibratoPulse = '\uE219'
	// Stockhausen irregular tremolo ("Morsen", like Morse code)
	StockhausenTremolo = '\uE232'
	// Bow behind bridge (sul ponticello)
	StringsBowBehindBridge = '\uE618'
	// Bow behind bridge on four strings
	StringsBowBehindBridgeFourStrings = '\uE62A'
	// Bow behind bridge on one string
	StringsBowBehindBridgeOneString = '\uE627'
	// Bow behind bridge on three strings
	StringsBowBehindBridgeThreeStrings = '\uE629'
	// Bow behind bridge on two strings
	StringsBowBehindBridgeTwoStrings = '\uE628'
	// Bow on top of bridge
	StringsBowOnBridge = '\uE619'
	// Bow on tailpiece
	StringsBowOnTailpiece = '\uE61A'
	// Change bow direction, indeterminate
	StringsChangeBowDirection = '\uE626'
	// Down bow
	StringsDownBow = '\uE610'
	// Down bow, away from body
	StringsDownBowAwayFromBody = '\uEE82'
	// Down bow, beyond bridge
	StringsDownBowBeyondBridge = '\uEE84'
	// Down bow, towards body
	StringsDownBowTowardsBody = '\uEE80'
	// Turned down bow
	StringsDownBowTurned = '\uE611'
	// Fouetté
	StringsFouette = '\uE622'
	// Half-harmonic
	StringsHalfHarmonic = '\uE615'
	// Harmonic
	StringsHarmonic = '\uE614'
	// Jeté (gettato) above
	StringsJeteAbove = '\uE620'
	// Jeté (gettato) below
	StringsJeteBelow = '\uE621'
	// Mute off
	StringsMuteOff = '\uE617'
	// Mute on
	StringsMuteOn = '\uE616'
	// Overpressure, down bow
	StringsOverpressureDownBow = '\uE61B'
	// Overpressure, no bow direction
	StringsOverpressureNoDirection = '\uE61F'
	// Overpressure possibile, down bow
	StringsOverpressurePossibileDownBow = '\uE61D'
	// Overpressure possibile, up bow
	StringsOverpressurePossibileUpBow = '\uE61E'
	// Overpressure, up bow
	StringsOverpressureUpBow = '\uE61C'
	// Scrape, circular clockwise
	StringsScrapeCircularClockwise = '\uEE88'
	// Scrape, circular counter-clockwise
	StringsScrapeCircularCounterclockwise = '\uEE89'
	// Scrape, parallel inward
	StringsScrapeParallelInward = '\uEE86'
	// Scrape, parallel outward
	StringsScrapeParallelOutward = '\uEE87'
	// Thumb position
	StringsThumbPosition = '\uE624'
	// Turned thumb position
	StringsThumbPositionTurned = '\uE625'
	// Triple chop, inward
	StringsTripleChopInward = '\uEE8A'
	// Triple chop, outward
	StringsTripleChopOutward = '\uEE8B'
	// Up bow
	StringsUpBow = '\uE612'
	// Up bow, away from body
	StringsUpBowAwayFromBody = '\uEE83'
	// Up bow, beyond bridge
	StringsUpBowBeyondBridge = '\uEE85'
	// Up bow, towards body
	StringsUpBowTowardsBody = '\uEE81'
	// Turned up bow
	StringsUpBowTurned = '\uE613'
	// Vibrato pulse accent (Saunders) for stem
	StringsVibratoPulse = '\uE623'
	// Swiss rudiments doublé black notehead
	SwissRudimentsNoteheadBlackDouble = '\uEE72'
	// Swiss rudiments flam black notehead
	SwissRudimentsNoteheadBlackFlam = '\uEE70'
	// Swiss rudiments doublé half (minim) notehead
	SwissRudimentsNoteheadHalfDouble = '\uEE73'
	// Swiss rudiments flam half (minim) notehead
	SwissRudimentsNoteheadHalfFlam = '\uEE71'
	// System divider
	SystemDivider = '\uE007'
	// Extra long system divider
	SystemDividerExtraLong = '\uE009'
	// Long system divider
	SystemDividerLong = '\uE008'
	// Augmentation dot
	TextAugmentationDot = '\uE1FC'
	// Black note, fractional 16th beam, long stem
	TextBlackNoteFrac16thLongStem = '\uE1F5'
	// Black note, fractional 16th beam, short stem
	TextBlackNoteFrac16thShortStem = '\uE1F4'
	// Black note, fractional 32nd beam, long stem
	TextBlackNoteFrac32ndLongStem = '\uE1F6'
	// Black note, fractional 8th beam, long stem
	TextBlackNoteFrac8thLongStem = '\uE1F3'
	// Black note, fractional 8th beam, short stem
	TextBlackNoteFrac8thShortStem = '\uE1F2'
	// Black note, long stem
	TextBlackNoteLongStem = '\uE1F1'
	// Black note, short stem
	TextBlackNoteShortStem = '\uE1F0'
	// Continuing 16th beam for long stem
	TextCont16thBeamLongStem = '\uE1FA'
	// Continuing 16th beam for short stem
	TextCont16thBeamShortStem = '\uE1F9'
	// Continuing 32nd beam for long stem
	TextCont32ndBeamLongStem = '\uE1FB'
	// Continuing 8th beam for long stem
	TextCont8thBeamLongStem = '\uE1F8'
	// Continuing 8th beam for short stem
	TextCont8thBeamShortStem = '\uE1F7'
	// Headless black note, fractional 16th beam, long stem
	TextHeadlessBlackNoteFrac16thLongStem = '\uE209'
	// Headless black note, fractional 16th beam, short stem
	TextHeadlessBlackNoteFrac16thShortStem = '\uE208'
	// Headless black note, fractional 32nd beam, long stem
	TextHeadlessBlackNoteFrac32ndLongStem = '\uE20A'
	// Headless black note, fractional 8th beam, long stem
	TextHeadlessBlackNoteFrac8thLongStem = '\uE207'
	// Headless black note, fractional 8th beam, short stem
	TextHeadlessBlackNoteFrac8thShortStem = '\uE206'
	// Headless black note, long stem
	TextHeadlessBlackNoteLongStem = '\uE205'
	// Headless black note, short stem
	TextHeadlessBlackNoteShortStem = '\uE204'
	// Tie
	TextTie = '\uE1FD'
	// Tuplet number 3 for long stem
	TextTuplet3LongStem = '\uE202'
	// Tuplet number 3 for short stem
	TextTuplet3ShortStem = '\uE1FF'
	// Tuplet bracket end for long stem
	TextTupletBracketEndLongStem = '\uE203'
	// Tuplet bracket end for short stem
	TextTupletBracketEndShortStem = '\uE200'
	// Tuplet bracket start for long stem
	TextTupletBracketStartLongStem = '\uE201'
	// Tuplet bracket start for short stem
	TextTupletBracketStartShortStem = '\uE1FE'
	// Time signature 0
	TimeSig0 = '\uE080'
	// Reversed time signature 0
	TimeSig0Reversed = '\uECF0'
	// Turned time signature 0
	TimeSig0Turned = '\uECE0'
	// Time signature 1
	TimeSig1 = '\uE081'
	// Reversed time signature 1
	TimeSig1Reversed = '\uECF1'
	// Turned time signature 1
	TimeSig1Turned = '\uECE1'
	// Time signature 2
	TimeSig2 = '\uE082'
	// Reversed time signature 2
	TimeSig2Reversed = '\uECF2'
	// Turned time signature 2
	TimeSig2Turned = '\uECE2'
	// Time signature 3
	TimeSig3 = '\uE083'
	// Reversed time signature 3
	TimeSig3Reversed = '\uECF3'
	// Turned time signature 3
	TimeSig3Turned = '\uECE3'
	// Time signature 4
	TimeSig4 = '\uE084'
	// Reversed time signature 4
	TimeSig4Reversed = '\uECF4'
	// Turned time signature 4
	TimeSig4Turned = '\uECE4'
	// Time signature 5
	TimeSig5 = '\uE085'
	// Reversed time signature 5
	TimeSig5Reversed = '\uECF5'
	// Turned time signature 5
	TimeSig5Turned = '\uECE5'
	// Time signature 6
	TimeSig6 = '\uE086'
	// Reversed time signature 6
	TimeSig6Reversed = '\uECF6'
	// Turned time signature 6
	TimeSig6Turned = '\uECE6'
	// Time signature 7
	TimeSig7 = '\uE087'
	// Reversed time signature 7
	TimeSig7Reversed = '\uECF7'
	// Turned time signature 7
	TimeSig7Turned = '\uECE7'
	// Time signature 8
	TimeSig8 = '\uE088'
	// Reversed time signature 8
	TimeSig8Reversed = '\uECF8'
	// Turned time signature 8
	TimeSig8Turned = '\uECE8'
	// Time signature 9
	TimeSig9 = '\uE089'
	// Reversed time signature 9
	TimeSig9Reversed = '\uECF9'
	// Turned time signature 9
	TimeSig9Turned = '\uECE9'
	// Left bracket for whole time signature
	TimeSigBracketLeft = '\uEC80'
	// Left bracket for numerator only
	TimeSigBracketLeftSmall = '\uEC82'
	// Right bracket for whole time signature
	TimeSigBracketRight = '\uEC81'
	// Right bracket for numerator only
	TimeSigBracketRightSmall = '\uEC83'
	// Control character for denominator digit
	TimeSigCombDenominator = '\uE09F'
	// Control character for numerator digit
	TimeSigCombNumerator = '\uE09E'
	// Time signature comma
	TimeSigComma = '\uE096'
	// Common time
	TimeSigCommon = '\uE08A'
	// Reversed common time
	TimeSigCommonReversed = '\uECFA'
	// Turned common time
	TimeSigCommonTurned = '\uECEA'
	// Cut time (Bach)
	TimeSigCut2 = '\uEC85'
	// Cut triple time (9/8)
	TimeSigCut3 = '\uEC86'
	// Cut time
	TimeSigCutCommon = '\uE08B'
	// Reversed cut time
	TimeSigCutCommonReversed = '\uECFB'
	// Turned cut time
	TimeSigCutCommonTurned = '\uECEB'
	// Time signature equals
	TimeSigEquals = '\uE08F'
	// Time signature fraction ½
	TimeSigFractionHalf = '\uE098'
	// Time signature fraction ⅓
	TimeSigFractionOneThird = '\uE09A'
	// Time signature fraction ¼
	TimeSigFractionQuarter = '\uE097'
	// Time signature fraction ¾
	TimeSigFractionThreeQuarters = '\uE099'
	// Time signature fraction ⅔
	TimeSigFractionTwoThirds = '\uE09B'
	// Time signature fraction slash
	TimeSigFractionalSlash = '\uE08E'
	// Time signature minus
	TimeSigMinus = '\uE090'
	// Time signature multiply
	TimeSigMultiply = '\uE091'
	// Open time signature (Penderecki)
	TimeSigOpenPenderecki = '\uE09D'
	// Left parenthesis for whole time signature
	TimeSigParensLeft = '\uE094'
	// Left parenthesis for numerator only
	TimeSigParensLeftSmall = '\uE092'
	// Right parenthesis for whole time signature
	TimeSigParensRight = '\uE095'
	// Right parenthesis for numerator only
	TimeSigParensRightSmall = '\uE093'
	// Time signature +
	TimeSigPlus = '\uE08C'
	// Time signature + (for numerators)
	TimeSigPlusSmall = '\uE08D'
	// Time signature slash separator
	TimeSigSlash = '\uEC84'
	// Open time signature
	TimeSigX = '\uE09C'
	// Combining tremolo 1
	Tremolo1 = '\uE220'
	// Combining tremolo 2
	Tremolo2 = '\uE221'
	// Combining tremolo 3
	Tremolo3 = '\uE222'
	// Combining tremolo 4
	Tremolo4 = '\uE223'
	// Combining tremolo 5
	Tremolo5 = '\uE224'
	// Divide measured tremolo by 2
	TremoloDivisiDots2 = '\uE22E'
	// Divide measured tremolo by 3
	TremoloDivisiDots3 = '\uE22F'
	// Divide measured tremolo by 4
	TremoloDivisiDots4 = '\uE230'
	// Divide measured tremolo by 6
	TremoloDivisiDots6 = '\uE231'
	// Fingered tremolo 1
	TremoloFingered1 = '\uE225'
	// Fingered tremolo 2
	TremoloFingered2 = '\uE226'
	// Fingered tremolo 3
	TremoloFingered3 = '\uE227'
	// Fingered tremolo 4
	TremoloFingered4 = '\uE228'
	// Fingered tremolo 5
	TremoloFingered5 = '\uE229'
	// Triple-tongue above
	TripleTongueAbove = '\uE5F2'
	// Triple-tongue below
	TripleTongueBelow = '\uE5F3'
	// Tuplet 0
	Tuplet0 = '\uE880'
	// Tuplet 1
	Tuplet1 = '\uE881'
	// Tuplet 2
	Tuplet2 = '\uE882'
	// Tuplet 3
	Tuplet3 = '\uE883'
	// Tuplet 4
	Tuplet4 = '\uE884'
	// Tuplet 5
	Tuplet5 = '\uE885'
	// Tuplet 6
	Tuplet6 = '\uE886'
	// Tuplet 7
	Tuplet7 = '\uE887'
	// Tuplet 8
	Tuplet8 = '\uE888'
	// Tuplet 9
	Tuplet9 = '\uE889'
	// Tuplet colon
	TupletColon = '\uE88A'
	// Wieniawski unmeasured tremolo
	UnmeasuredTremolo = '\uE22C'
	// Wieniawski unmeasured tremolo (simpler)
	UnmeasuredTremoloSimple = '\uE22D'
	// Unpitched percussion clef 1
	UnpitchedPercussionClef1 = '\uE069'
	// Unpitched percussion clef 2
	UnpitchedPercussionClef2 = '\uE06A'
	// Ventiduesima
	Ventiduesima = '\uE517'
	// Ventiduesima alta
	VentiduesimaAlta = '\uE518'
	// Ventiduesima bassa
	VentiduesimaBassa = '\uE519'
	// Ventiduesima bassa (mb)
	VentiduesimaBassaMb = '\uE51E'
	// Finger click (Stockhausen)
	VocalFingerClickStockhausen = '\uE649'
	// Halb gesungen (semi-sprechgesang)
	VocalHalbGesungen = '\uE64B'
	// Mouth closed
	VocalMouthClosed = '\uE640'
	// Mouth open
	VocalMouthOpen = '\uE642'
	// Mouth pursed
	VocalMouthPursed = '\uE644'
	// Mouth slightly open
	VocalMouthSlightlyOpen = '\uE641'
	// Mouth wide open
	VocalMouthWideOpen = '\uE643'
	// Nasal voice
	VocalNasalVoice = '\uE647'
	// Sprechgesang
	VocalSprechgesang = '\uE645'
	// Tongue click (Stockhausen)
	VocalTongueClickStockhausen = '\uE648'
	// Tongue and finger click (Stockhausen)
	VocalTongueFingerClickStockhausen = '\uE64A'
	// Combining sussurando for stem
	VocalsSussurando = '\uE646'
	// Arpeggiato wiggle segment, downwards
	WiggleArpeggiatoDown = '\uEAAA'
	// Arpeggiato arrowhead down
	WiggleArpeggiatoDownArrow = '\uEAAE'
	// Arpeggiato downward swash
	WiggleArpeggiatoDownSwash = '\uEAAC'
	// Arpeggiato wiggle segment, upwards
	WiggleArpeggiatoUp = '\uEAA9'
	// Arpeggiato arrowhead up
	WiggleArpeggiatoUpArrow = '\uEAAD'
	// Arpeggiato upward swash
	WiggleArpeggiatoUpSwash = '\uEAAB'
	// Circular motion segment
	WiggleCircular = '\uEAC9'
	// Constant circular motion segment
	WiggleCircularConstant = '\uEAC0'
	// Constant circular motion segment (flipped)
	WiggleCircularConstantFlipped = '\uEAC1'
	// Constant circular motion segment (flipped, large)
	WiggleCircularConstantFlippedLarge = '\uEAC3'
	// Constant circular motion segment (large)
	WiggleCircularConstantLarge = '\uEAC2'
	// Circular motion end
	WiggleCircularEnd = '\uEACB'
	// Circular motion segment, large
	WiggleCircularLarge = '\uEAC8'
	// Circular motion segment, larger
	WiggleCircularLarger = '\uEAC7'
	// Circular motion segment, larger still
	WiggleCircularLargerStill = '\uEAC6'
	// Circular motion segment, largest
	WiggleCircularLargest = '\uEAC5'
	// Circular motion segment, small
	WiggleCircularSmall = '\uEACA'
	// Circular motion start
	WiggleCircularStart = '\uEAC4'
	// Glissando wiggle segment
	WiggleGlissando = '\uEAAF'
	// Group glissando 1
	WiggleGlissandoGroup1 = '\uEABD'
	// Group glissando 2
	WiggleGlissandoGroup2 = '\uEABE'
	// Group glissando 3
	WiggleGlissandoGroup3 = '\uEABF'
	// Quasi-random squiggle 1
	WiggleRandom1 = '\uEAF0'
	// Quasi-random squiggle 2
	WiggleRandom2 = '\uEAF1'
	// Quasi-random squiggle 3
	WiggleRandom3 = '\uEAF2'
	// Quasi-random squiggle 4
	WiggleRandom4 = '\uEAF3'
	// Sawtooth line segment
	WiggleSawtooth = '\uEABB'
	// Narrow sawtooth line segment
	WiggleSawtoothNarrow = '\uEABA'
	// Wide sawtooth line segment
	WiggleSawtoothWide = '\uEABC'
	// Square wave line segment
	WiggleSquareWave = '\uEAB8'
	// Narrow square wave line segment
	WiggleSquareWaveNarrow = '\uEAB7'
	// Wide square wave line segment
	WiggleSquareWaveWide = '\uEAB9'
	// Trill wiggle segment
	WiggleTrill = '\uEAA4'
	// Trill wiggle segment, fast
	WiggleTrillFast = '\uEAA3'
	// Trill wiggle segment, faster
	WiggleTrillFaster = '\uEAA2'
	// Trill wiggle segment, faster still
	WiggleTrillFasterStill = '\uEAA1'
	// Trill wiggle segment, fastest
	WiggleTrillFastest = '\uEAA0'
	// Trill wiggle segment, slow
	WiggleTrillSlow = '\uEAA5'
	// Trill wiggle segment, slower
	WiggleTrillSlower = '\uEAA6'
	// Trill wiggle segment, slower still
	WiggleTrillSlowerStill = '\uEAA7'
	// Trill wiggle segment, slowest
	WiggleTrillSlowest = '\uEAA8'
	// Vibrato largest, slower
	WiggleVIbratoLargestSlower = '\uEAEE'
	// Vibrato medium, slower
	WiggleVIbratoMediumSlower = '\uEAE0'
	// Vibrato / shake wiggle segment
	WiggleVibrato = '\uEAB0'
	// Vibrato large, fast
	WiggleVibratoLargeFast = '\uEAE5'
	// Vibrato large, faster
	WiggleVibratoLargeFaster = '\uEAE4'
	// Vibrato large, faster still
	WiggleVibratoLargeFasterStill = '\uEAE3'
	// Vibrato large, fastest
	WiggleVibratoLargeFastest = '\uEAE2'
	// Vibrato large, slow
	WiggleVibratoLargeSlow = '\uEAE6'
	// Vibrato large, slower
	WiggleVibratoLargeSlower = '\uEAE7'
	// Vibrato large, slowest
	WiggleVibratoLargeSlowest = '\uEAE8'
	// Vibrato largest, fast
	WiggleVibratoLargestFast = '\uEAEC'
	// Vibrato largest, faster
	WiggleVibratoLargestFaster = '\uEAEB'
	// Vibrato largest, faster still
	WiggleVibratoLargestFasterStill = '\uEAEA'
	// Vibrato largest, fastest
	WiggleVibratoLargestFastest = '\uEAE9'
	// Vibrato largest, slow
	WiggleVibratoLargestSlow = '\uEAED'
	// Vibrato largest, slowest
	WiggleVibratoLargestSlowest = '\uEAEF'
	// Vibrato medium, fast
	WiggleVibratoMediumFast = '\uEADE'
	// Vibrato medium, faster
	WiggleVibratoMediumFaster = '\uEADD'
	// Vibrato medium, faster still
	WiggleVibratoMediumFasterStill = '\uEADC'
	// Vibrato medium, fastest
	WiggleVibratoMediumFastest = '\uEADB'
	// Vibrato medium, slow
	WiggleVibratoMediumSlow = '\uEADF'
	// Vibrato medium, slowest
	WiggleVibratoMediumSlowest = '\uEAE1'
	// Vibrato small, fast
	WiggleVibratoSmallFast = '\uEAD7'
	// Vibrato small, faster
	WiggleVibratoSmallFaster = '\uEAD6'
	// Vibrato small, faster still
	WiggleVibratoSmallFasterStill = '\uEAD5'
	// Vibrato small, fastest
	WiggleVibratoSmallFastest = '\uEAD4'
	// Vibrato small, slow
	WiggleVibratoSmallSlow = '\uEAD8'
	// Vibrato small, slower
	WiggleVibratoSmallSlower = '\uEAD9'
	// Vibrato small, slowest
	WiggleVibratoSmallSlowest = '\uEADA'
	// Vibrato smallest, fast
	WiggleVibratoSmallestFast = '\uEAD0'
	// Vibrato smallest, faster
	WiggleVibratoSmallestFaster = '\uEACF'
	// Vibrato smallest, faster still
	WiggleVibratoSmallestFasterStill = '\uEACE'
	// Vibrato smallest, fastest
	WiggleVibratoSmallestFastest = '\uEACD'
	// Vibrato smallest, slow
	WiggleVibratoSmallestSlow = '\uEAD1'
	// Vibrato smallest, slower
	WiggleVibratoSmallestSlower = '\uEAD2'
	// Vibrato smallest, slowest
	WiggleVibratoSmallestSlowest = '\uEAD3'
	// Vibrato start
	WiggleVibratoStart = '\uEACC'
	// Wide vibrato / shake wiggle segment
	WiggleVibratoWide = '\uEAB1'
	// Wavy line segment
	WiggleWavy = '\uEAB5'
	// Narrow wavy line segment
	WiggleWavyNarrow = '\uEAB4'
	// Wide wavy line segment
	WiggleWavyWide = '\uEAB6'
	// Closed hole
	WindClosedHole = '\uE5F4'
	// Flatter embouchure
	WindFlatEmbouchure = '\uE5FB'
	// Half-closed hole
	WindHalfClosedHole1 = '\uE5F6'
	// Half-closed hole 2
	WindHalfClosedHole2 = '\uE5F7'
	// Half-open hole
	WindHalfClosedHole3 = '\uE5F8'
	// Somewhat relaxed embouchure
	WindLessRelaxedEmbouchure = '\uE5FE'
	// Somewhat tight embouchure
	WindLessTightEmbouchure = '\uE600'
	// Mouthpiece or hand pop
	WindMouthpiecePop = '\uE60A'
	// Combining multiphonics (black) for stem
	WindMultiphonicsBlackStem = '\uE607'
	// Combining multiphonics (black and white) for stem
	WindMultiphonicsBlackWhiteStem = '\uE609'
	// Combining multiphonics (white) for stem
	WindMultiphonicsWhiteStem = '\uE608'
	// Open hole
	WindOpenHole = '\uE5F9'
	// Much more reed (push inwards)
	WindReedPositionIn = '\uE606'
	// Normal reed position
	WindReedPositionNormal = '\uE604'
	// Very little reed (pull outwards)
	WindReedPositionOut = '\uE605'
	// Relaxed embouchure
	WindRelaxedEmbouchure = '\uE5FD'
	// Rim only
	WindRimOnly = '\uE60B'
	// Sharper embouchure
	WindSharpEmbouchure = '\uE5FC'
	// Very tight embouchure / strong air pressure
	WindStrongAirPressure = '\uE603'
	// Three-quarters closed hole
	WindThreeQuartersClosedHole = '\uE5F5'
	// Tight embouchure
	WindTightEmbouchure = '\uE5FF'
	// Trill key
	WindTrillKey = '\uE5FA'
	// Very tight embouchure
	WindVeryTightEmbouchure = '\uE601'
	// Very relaxed embouchure / weak air-pressure
	WindWeakAirPressure = '\uE602'
)

var nameToCodepoint = map[string]rune{
	"4stringTabClef": FourstringTabClef,
	"6stringTabClef": SixstringTabClef,
	"accSagittal11LargeDiesisDown": AccSagittal11LargeDiesisDown,
	"accSagittal11LargeDiesisUp": AccSagittal11LargeDiesisUp,
	"accSagittal11MediumDiesisDown": AccSagittal11MediumDiesisDown,
	"accSagittal11MediumDiesisUp": AccSagittal11MediumDiesisUp,
	"accSagittal11v19LargeDiesisDown": AccSagittal11v19LargeDiesisDown,
	"accSagittal11v19LargeDiesisUp": AccSagittal11v19LargeDiesisUp,
	"accSagittal11v19MediumDiesisDown": AccSagittal11v19MediumDiesisDown,
	"accSagittal11v19MediumDiesisUp": AccSagittal11v19MediumDiesisUp,
	"accSagittal11v49CommaDown": AccSagittal11v49CommaDown,
	"accSagittal11v49CommaUp": AccSagittal11v49CommaUp,
	"accSagittal143CommaDown": AccSagittal143CommaDown,
	"accSagittal143CommaUp": AccSagittal143CommaUp,
	"accSagittal17CommaDown": AccSagittal17CommaDown,
	"accSagittal17CommaUp": AccSagittal17CommaUp,
	"accSagittal17KleismaDown": AccSagittal17KleismaDown,
	"accSagittal17KleismaUp": AccSagittal17KleismaUp,
	"accSagittal19CommaDown": AccSagittal19CommaDown,
	"accSagittal19CommaUp": AccSagittal19CommaUp,
	"accSagittal19SchismaDown": AccSagittal19SchismaDown,
	"accSagittal19SchismaUp": AccSagittal19SchismaUp,
	"accSagittal1MinaDown": AccSagittal1MinaDown,
	"accSagittal1MinaUp": AccSagittal1MinaUp,
	"accSagittal1TinaDown": AccSagittal1TinaDown,
	"accSagittal1TinaUp": AccSagittal1TinaUp,
	"accSagittal23CommaDown": AccSagittal23CommaDown,
	"accSagittal23CommaUp": AccSagittal23CommaUp,
	"accSagittal23SmallDiesisDown": AccSagittal23SmallDiesisDown,
	"accSagittal23SmallDiesisUp": AccSagittal23SmallDiesisUp,
	"accSagittal25SmallDiesisDown": AccSagittal25SmallDiesisDown,
	"accSagittal25SmallDiesisUp": AccSagittal25SmallDiesisUp,
	"accSagittal2MinasDown": AccSagittal2MinasDown,
	"accSagittal2MinasUp": AccSagittal2MinasUp,
	"accSagittal2TinasDown": AccSagittal2TinasDown,
	"accSagittal2TinasUp": AccSagittal2TinasUp,
	"accSagittal35LargeDiesisDown": AccSagittal35LargeDiesisDown,
	"accSagittal35LargeDiesisUp": AccSagittal35LargeDiesisUp,
	"accSagittal35MediumDiesisDown": AccSagittal35MediumDiesisDown,
	"accSagittal35MediumDiesisUp": AccSagittal35MediumDiesisUp,
	"accSagittal3TinasDown": AccSagittal3TinasDown,
	"accSagittal3TinasUp": AccSagittal3TinasUp,
	"accSagittal49LargeDiesisDown": AccSagittal49LargeDiesisDown,
	"accSagittal49LargeDiesisUp": AccSagittal49LargeDiesisUp,
	"accSagittal49MediumDiesisDown": AccSagittal49MediumDiesisDown,
	"accSagittal49MediumDiesisUp": AccSagittal49MediumDiesisUp,
	"accSagittal49SmallDiesisDown": AccSagittal49SmallDiesisDown,
	"accSagittal49SmallDiesisUp": AccSagittal49SmallDiesisUp,
	"accSagittal4TinasDown": AccSagittal4TinasDown,
	"accSagittal4TinasUp": AccSagittal4TinasUp,
	"accSagittal55CommaDown": AccSagittal55CommaDown,
	"accSagittal55CommaUp": AccSagittal55CommaUp,
	"accSagittal5CommaDown": AccSagittal5CommaDown,
	"accSagittal5CommaUp": AccSagittal5CommaUp,
	"accSagittal5TinasDown": AccSagittal5TinasDown,
	"accSagittal5TinasUp": AccSagittal5TinasUp,
	"accSagittal5v11SmallDiesisDown": AccSagittal5v11SmallDiesisDown,
	"accSagittal5v11SmallDiesisUp": AccSagittal5v11SmallDiesisUp,
	"accSagittal5v13LargeDiesisDown": AccSagittal5v13LargeDiesisDown,
	"accSagittal5v13LargeDiesisUp": AccSagittal5v13LargeDiesisUp,
	"accSagittal5v13MediumDiesisDown": AccSagittal5v13MediumDiesisDown,
	"accSagittal5v13MediumDiesisUp": AccSagittal5v13MediumDiesisUp,
	"accSagittal5v19CommaDown": AccSagittal5v19CommaDown,
	"accSagittal5v19CommaUp": AccSagittal5v19CommaUp,
	"accSagittal5v23SmallDiesisDown": AccSagittal5v23SmallDiesisDown,
	"accSagittal5v23SmallDiesisUp": AccSagittal5v23SmallDiesisUp,
	"accSagittal5v49MediumDiesisDown": AccSagittal5v49MediumDiesisDown,
	"accSagittal5v49MediumDiesisUp": AccSagittal5v49MediumDiesisUp,
	"accSagittal5v7KleismaDown": AccSagittal5v7KleismaDown,
	"accSagittal5v7KleismaUp": AccSagittal5v7KleismaUp,
	"accSagittal6TinasDown": AccSagittal6TinasDown,
	"accSagittal6TinasUp": AccSagittal6TinasUp,
	"accSagittal7CommaDown": AccSagittal7CommaDown,
	"accSagittal7CommaUp": AccSagittal7CommaUp,
	"accSagittal7TinasDown": AccSagittal7TinasDown,
	"accSagittal7TinasUp": AccSagittal7TinasUp,
	"accSagittal7v11CommaDown": AccSagittal7v11CommaDown,
	"accSagittal7v11CommaUp": AccSagittal7v11CommaUp,
	"accSagittal7v11KleismaDown": AccSagittal7v11KleismaDown,
	"accSagittal7v11KleismaUp": AccSagittal7v11KleismaUp,
	"accSagittal7v19CommaDown": AccSagittal7v19CommaDown,
	"accSagittal7v19CommaUp": AccSagittal7v19CommaUp,
	"accSagittal8TinasDown": AccSagittal8TinasDown,
	"accSagittal8TinasUp": AccSagittal8TinasUp,
	"accSagittal9TinasDown": AccSagittal9TinasDown,
	"accSagittal9TinasUp": AccSagittal9TinasUp,
	"accSagittalAcute": AccSagittalAcute,
	"accSagittalDoubleFlat": AccSagittalDoubleFlat,
	"accSagittalDoubleFlat11v49CUp": AccSagittalDoubleFlat11v49CUp,
	"accSagittalDoubleFlat143CUp": AccSagittalDoubleFlat143CUp,
	"accSagittalDoubleFlat17CUp": AccSagittalDoubleFlat17CUp,
	"accSagittalDoubleFlat17kUp": AccSagittalDoubleFlat17kUp,
	"accSagittalDoubleFlat19CUp": AccSagittalDoubleFlat19CUp,
	"accSagittalDoubleFlat19sUp": AccSagittalDoubleFlat19sUp,
	"accSagittalDoubleFlat23CUp": AccSagittalDoubleFlat23CUp,
	"accSagittalDoubleFlat23SUp": AccSagittalDoubleFlat23SUp,
	"accSagittalDoubleFlat25SUp": AccSagittalDoubleFlat25SUp,
	"accSagittalDoubleFlat49SUp": AccSagittalDoubleFlat49SUp,
	"accSagittalDoubleFlat55CUp": AccSagittalDoubleFlat55CUp,
	"accSagittalDoubleFlat5CUp": AccSagittalDoubleFlat5CUp,
	"accSagittalDoubleFlat5v11SUp": AccSagittalDoubleFlat5v11SUp,
	"accSagittalDoubleFlat5v19CUp": AccSagittalDoubleFlat5v19CUp,
	"accSagittalDoubleFlat5v23SUp": AccSagittalDoubleFlat5v23SUp,
	"accSagittalDoubleFlat5v7kUp": AccSagittalDoubleFlat5v7kUp,
	"accSagittalDoubleFlat7CUp": AccSagittalDoubleFlat7CUp,
	"accSagittalDoubleFlat7v11CUp": AccSagittalDoubleFlat7v11CUp,
	"accSagittalDoubleFlat7v11kUp": AccSagittalDoubleFlat7v11kUp,
	"accSagittalDoubleFlat7v19CUp": AccSagittalDoubleFlat7v19CUp,
	"accSagittalDoubleSharp": AccSagittalDoubleSharp,
	"accSagittalDoubleSharp11v49CDown": AccSagittalDoubleSharp11v49CDown,
	"accSagittalDoubleSharp143CDown": AccSagittalDoubleSharp143CDown,
	"accSagittalDoubleSharp17CDown": AccSagittalDoubleSharp17CDown,
	"accSagittalDoubleSharp17kDown": AccSagittalDoubleSharp17kDown,
	"accSagittalDoubleSharp19CDown": AccSagittalDoubleSharp19CDown,
	"accSagittalDoubleSharp19sDown": AccSagittalDoubleSharp19sDown,
	"accSagittalDoubleSharp23CDown": AccSagittalDoubleSharp23CDown,
	"accSagittalDoubleSharp23SDown": AccSagittalDoubleSharp23SDown,
	"accSagittalDoubleSharp25SDown": AccSagittalDoubleSharp25SDown,
	"accSagittalDoubleSharp49SDown": AccSagittalDoubleSharp49SDown,
	"accSagittalDoubleSharp55CDown": AccSagittalDoubleSharp55CDown,
	"accSagittalDoubleSharp5CDown": AccSagittalDoubleSharp5CDown,
	"accSagittalDoubleSharp5v11SDown": AccSagittalDoubleSharp5v11SDown,
	"accSagittalDoubleSharp5v19CDown": AccSagittalDoubleSharp5v19CDown,
	"accSagittalDoubleSharp5v23SDown": AccSagittalDoubleSharp5v23SDown,
	"accSagittalDoubleSharp5v7kDown": AccSagittalDoubleSharp5v7kDown,
	"accSagittalDoubleSharp7CDown": AccSagittalDoubleSharp7CDown,
	"accSagittalDoubleSharp7v11CDown": AccSagittalDoubleSharp7v11CDown,
	"accSagittalDoubleSharp7v11kDown": AccSagittalDoubleSharp7v11kDown,
	"accSagittalDoubleSharp7v19CDown": AccSagittalDoubleSharp7v19CDown,
	"accSagittalFlat": AccSagittalFlat,
	"accSagittalFlat11LDown": AccSagittalFlat11LDown,
	"accSagittalFlat11MDown": AccSagittalFlat11MDown,
	"accSagittalFlat11v19LDown": AccSagittalFlat11v19LDown,
	"accSagittalFlat11v19MDown": AccSagittalFlat11v19MDown,
	"accSagittalFlat11v49CDown": AccSagittalFlat11v49CDown,
	"accSagittalFlat11v49CUp": AccSagittalFlat11v49CUp,
	"accSagittalFlat143CDown": AccSagittalFlat143CDown,
	"accSagittalFlat143CUp": AccSagittalFlat143CUp,
	"accSagittalFlat17CDown": AccSagittalFlat17CDown,
	"accSagittalFlat17CUp": AccSagittalFlat17CUp,
	"accSagittalFlat17kDown": AccSagittalFlat17kDown,
	"accSagittalFlat17kUp": AccSagittalFlat17kUp,
	"accSagittalFlat19CDown": AccSagittalFlat19CDown,
	"accSagittalFlat19CUp": AccSagittalFlat19CUp,
	"accSagittalFlat19sDown": AccSagittalFlat19sDown,
	"accSagittalFlat19sUp": AccSagittalFlat19sUp,
	"accSagittalFlat23CDown": AccSagittalFlat23CDown,
	"accSagittalFlat23CUp": AccSagittalFlat23CUp,
	"accSagittalFlat23SDown": AccSagittalFlat23SDown,
	"accSagittalFlat23SUp": AccSagittalFlat23SUp,
	"accSagittalFlat25SDown": AccSagittalFlat25SDown,
	"accSagittalFlat25SUp": AccSagittalFlat25SUp,
	"accSagittalFlat35LDown": AccSagittalFlat35LDown,
	"accSagittalFlat35MDown": AccSagittalFlat35MDown,
	"accSagittalFlat49LDown": AccSagittalFlat49LDown,
	"accSagittalFlat49MDown": AccSagittalFlat49MDown,
	"accSagittalFlat49SDown": AccSagittalFlat49SDown,
	"accSagittalFlat49SUp": AccSagittalFlat49SUp,
	"accSagittalFlat55CDown": AccSagittalFlat55CDown,
	"accSagittalFlat55CUp": AccSagittalFlat55CUp,
	"accSagittalFlat5CDown": AccSagittalFlat5CDown,
	"accSagittalFlat5CUp": AccSagittalFlat5CUp,
	"accSagittalFlat5v11SDown": AccSagittalFlat5v11SDown,
	"accSagittalFlat5v11SUp": AccSagittalFlat5v11SUp,
	"accSagittalFlat5v13LDown": AccSagittalFlat5v13LDown,
	"accSagittalFlat5v13MDown": AccSagittalFlat5v13MDown,
	"accSagittalFlat5v19CDown": AccSagittalFlat5v19CDown,
	"accSagittalFlat5v19CUp": AccSagittalFlat5v19CUp,
	"accSagittalFlat5v23SDown": AccSagittalFlat5v23SDown,
	"accSagittalFlat5v23SUp": AccSagittalFlat5v23SUp,
	"accSagittalFlat5v49MDown": AccSagittalFlat5v49MDown,
	"accSagittalFlat5v7kDown": AccSagittalFlat5v7kDown,
	"accSagittalFlat5v7kUp": AccSagittalFlat5v7kUp,
	"accSagittalFlat7CDown": AccSagittalFlat7CDown,
	"accSagittalFlat7CUp": AccSagittalFlat7CUp,
	"accSagittalFlat7v11CDown": AccSagittalFlat7v11CDown,
	"accSagittalFlat7v11CUp": AccSagittalFlat7v11CUp,
	"accSagittalFlat7v11kDown": AccSagittalFlat7v11kDown,
	"accSagittalFlat7v11kUp": AccSagittalFlat7v11kUp,
	"accSagittalFlat7v19CDown": AccSagittalFlat7v19CDown,
	"accSagittalFlat7v19CUp": AccSagittalFlat7v19CUp,
	"accSagittalFractionalTinaDown": AccSagittalFractionalTinaDown,
	"accSagittalFractionalTinaUp": AccSagittalFractionalTinaUp,
	"accSagittalGrave": AccSagittalGrave,
	"accSagittalShaftDown": AccSagittalShaftDown,
	"accSagittalShaftUp": AccSagittalShaftUp,
	"accSagittalSharp": AccSagittalSharp,
	"accSagittalSharp11LUp": AccSagittalSharp11LUp,
	"accSagittalSharp11MUp": AccSagittalSharp11MUp,
	"accSagittalSharp11v19LUp": AccSagittalSharp11v19LUp,
	"accSagittalSharp11v19MUp": AccSagittalSharp11v19MUp,
	"accSagittalSharp11v49CDown": AccSagittalSharp11v49CDown,
	"accSagittalSharp11v49CUp": AccSagittalSharp11v49CUp,
	"accSagittalSharp143CDown": AccSagittalSharp143CDown,
	"accSagittalSharp143CUp": AccSagittalSharp143CUp,
	"accSagittalSharp17CDown": AccSagittalSharp17CDown,
	"accSagittalSharp17CUp": AccSagittalSharp17CUp,
	"accSagittalSharp17kDown": AccSagittalSharp17kDown,
	"accSagittalSharp17kUp": AccSagittalSharp17kUp,
	"accSagittalSharp19CDown": AccSagittalSharp19CDown,
	"accSagittalSharp19CUp": AccSagittalSharp19CUp,
	"accSagittalSharp19sDown": AccSagittalSharp19sDown,
	"accSagittalSharp19sUp": AccSagittalSharp19sUp,
	"accSagittalSharp23CDown": AccSagittalSharp23CDown,
	"accSagittalSharp23CUp": AccSagittalSharp23CUp,
	"accSagittalSharp23SDown": AccSagittalSharp23SDown,
	"accSagittalSharp23SUp": AccSagittalSharp23SUp,
	"accSagittalSharp25SDown": AccSagittalSharp25SDown,
	"accSagittalSharp25SUp": AccSagittalSharp25SUp,
	"accSagittalSharp35LUp": AccSagittalSharp35LUp,
	"accSagittalSharp35MUp": AccSagittalSharp35MUp,
	"accSagittalSharp49LUp": AccSagittalSharp49LUp,
	"accSagittalSharp49MUp": AccSagittalSharp49MUp,
	"accSagittalSharp49SDown": AccSagittalSharp49SDown,
	"accSagittalSharp49SUp": AccSagittalSharp49SUp,
	"accSagittalSharp55CDown": AccSagittalSharp55CDown,
	"accSagittalSharp55CUp": AccSagittalSharp55CUp,
	"accSagittalSharp5CDown": AccSagittalSharp5CDown,
	"accSagittalSharp5CUp": AccSagittalSharp5CUp,
	"accSagittalSharp5v11SDown": AccSagittalSharp5v11SDown,
	"accSagittalSharp5v11SUp": AccSagittalSharp5v11SUp,
	"accSagittalSharp5v13LUp": AccSagittalSharp5v13LUp,
	"accSagittalSharp5v13MUp": AccSagittalSharp5v13MUp,
	"accSagittalSharp5v19CDown": AccSagittalSharp5v19CDown,
	"accSagittalSharp5v19CUp": AccSagittalSharp5v19CUp,
	"accSagittalSharp5v23SDown": AccSagittalSharp5v23SDown,
	"accSagittalSharp5v23SUp": AccSagittalSharp5v23SUp,
	"accSagittalSharp5v49MUp": AccSagittalSharp5v49MUp,
	"accSagittalSharp5v7kDown": AccSagittalSharp5v7kDown,
	"accSagittalSharp5v7kUp": AccSagittalSharp5v7kUp,
	"accSagittalSharp7CDown": AccSagittalSharp7CDown,
	"accSagittalSharp7CUp": AccSagittalSharp7CUp,
	"accSagittalSharp7v11CDown": AccSagittalSharp7v11CDown,
	"accSagittalSharp7v11CUp": AccSagittalSharp7v11CUp,
	"accSagittalSharp7v11kDown": AccSagittalSharp7v11kDown,
	"accSagittalSharp7v11kUp": AccSagittalSharp7v11kUp,
	"accSagittalSharp7v19CDown": AccSagittalSharp7v19CDown,
	"accSagittalSharp7v19CUp": AccSagittalSharp7v19CUp,
	"accSagittalUnused1": AccSagittalUnused1,
	"accSagittalUnused2": AccSagittalUnused2,
	"accSagittalUnused3": AccSagittalUnused3,
	"accSagittalUnused4": AccSagittalUnused4,
	"accdnCombDot": AccdnCombDot,
	"accdnCombLH2RanksEmpty": AccdnCombLH2RanksEmpty,
	"accdnCombLH3RanksEmptySquare": AccdnCombLH3RanksEmptySquare,
	"accdnCombRH3RanksEmpty": AccdnCombRH3RanksEmpty,
	"accdnCombRH4RanksEmpty": AccdnCombRH4RanksEmpty,
	"accdnDiatonicClef": AccdnDiatonicClef,
	"accdnLH2Ranks16Round": AccdnLH2Ranks16Round,
	"accdnLH2Ranks8Plus16Round": AccdnLH2Ranks8Plus16Round,
	"accdnLH2Ranks8Round": AccdnLH2Ranks8Round,
	"accdnLH2RanksFullMasterRound": AccdnLH2RanksFullMasterRound,
	"accdnLH2RanksMasterPlus16Round": AccdnLH2RanksMasterPlus16Round,
	"accdnLH2RanksMasterRound": AccdnLH2RanksMasterRound,
	"accdnLH3Ranks2Plus8Square": AccdnLH3Ranks2Plus8Square,
	"accdnLH3Ranks2Square": AccdnLH3Ranks2Square,
	"accdnLH3Ranks8Square": AccdnLH3Ranks8Square,
	"accdnLH3RanksDouble8Square": AccdnLH3RanksDouble8Square,
	"accdnLH3RanksTuttiSquare": AccdnLH3RanksTuttiSquare,
	"accdnPull": AccdnPull,
	"accdnPush": AccdnPush,
	"accdnRH3RanksAccordion": AccdnRH3RanksAccordion,
	"accdnRH3RanksAuthenticMusette": AccdnRH3RanksAuthenticMusette,
	"accdnRH3RanksBandoneon": AccdnRH3RanksBandoneon,
	"accdnRH3RanksBassoon": AccdnRH3RanksBassoon,
	"accdnRH3RanksClarinet": AccdnRH3RanksClarinet,
	"accdnRH3RanksDoubleTremoloLower8ve": AccdnRH3RanksDoubleTremoloLower8ve,
	"accdnRH3RanksDoubleTremoloUpper8ve": AccdnRH3RanksDoubleTremoloUpper8ve,
	"accdnRH3RanksFullFactory": AccdnRH3RanksFullFactory,
	"accdnRH3RanksHarmonium": AccdnRH3RanksHarmonium,
	"accdnRH3RanksImitationMusette": AccdnRH3RanksImitationMusette,
	"accdnRH3RanksLowerTremolo8": AccdnRH3RanksLowerTremolo8,
	"accdnRH3RanksMaster": AccdnRH3RanksMaster,
	"accdnRH3RanksOboe": AccdnRH3RanksOboe,
	"accdnRH3RanksOrgan": AccdnRH3RanksOrgan,
	"accdnRH3RanksPiccolo": AccdnRH3RanksPiccolo,
	"accdnRH3RanksTremoloLower8ve": AccdnRH3RanksTremoloLower8ve,
	"accdnRH3RanksTremoloUpper8ve": AccdnRH3RanksTremoloUpper8ve,
	"accdnRH3RanksTwoChoirs": AccdnRH3RanksTwoChoirs,
	"accdnRH3RanksUpperTremolo8": AccdnRH3RanksUpperTremolo8,
	"accdnRH3RanksViolin": AccdnRH3RanksViolin,
	"accdnRH4RanksAlto": AccdnRH4RanksAlto,
	"accdnRH4RanksBassAlto": AccdnRH4RanksBassAlto,
	"accdnRH4RanksMaster": AccdnRH4RanksMaster,
	"accdnRH4RanksSoftBass": AccdnRH4RanksSoftBass,
	"accdnRH4RanksSoftTenor": AccdnRH4RanksSoftTenor,
	"accdnRH4RanksSoprano": AccdnRH4RanksSoprano,
	"accdnRH4RanksTenor": AccdnRH4RanksTenor,
	"accdnRicochet2": AccdnRicochet2,
	"accdnRicochet3": AccdnRicochet3,
	"accdnRicochet4": AccdnRicochet4,
	"accdnRicochet5": AccdnRicochet5,
	"accdnRicochet6": AccdnRicochet6,
	"accdnRicochetStem2": AccdnRicochetStem2,
	"accdnRicochetStem3": AccdnRicochetStem3,
	"accdnRicochetStem4": AccdnRicochetStem4,
	"accdnRicochetStem5": AccdnRicochetStem5,
	"accdnRicochetStem6": AccdnRicochetStem6,
	"accidental1CommaFlat": Accidental1CommaFlat,
	"accidental1CommaSharp": Accidental1CommaSharp,
	"accidental2CommaFlat": Accidental2CommaFlat,
	"accidental2CommaSharp": Accidental2CommaSharp,
	"accidental3CommaFlat": Accidental3CommaFlat,
	"accidental3CommaSharp": Accidental3CommaSharp,
	"accidental4CommaFlat": Accidental4CommaFlat,
	"accidental5CommaSharp": Accidental5CommaSharp,
	"accidentalArrowDown": AccidentalArrowDown,
	"accidentalArrowUp": AccidentalArrowUp,
	"accidentalBakiyeFlat": AccidentalBakiyeFlat,
	"accidentalBakiyeSharp": AccidentalBakiyeSharp,
	"accidentalBracketLeft": AccidentalBracketLeft,
	"accidentalBracketRight": AccidentalBracketRight,
	"accidentalBuyukMucennebFlat": AccidentalBuyukMucennebFlat,
	"accidentalBuyukMucennebSharp": AccidentalBuyukMucennebSharp,
	"accidentalCombiningCloseCurlyBrace": AccidentalCombiningCloseCurlyBrace,
	"accidentalCombiningLower17Schisma": AccidentalCombiningLower17Schisma,
	"accidentalCombiningLower19Schisma": AccidentalCombiningLower19Schisma,
	"accidentalCombiningLower23Limit29LimitComma": AccidentalCombiningLower23Limit29LimitComma,
	"accidentalCombiningLower29LimitComma": AccidentalCombiningLower29LimitComma,
	"accidentalCombiningLower31Schisma": AccidentalCombiningLower31Schisma,
	"accidentalCombiningLower37Quartertone": AccidentalCombiningLower37Quartertone,
	"accidentalCombiningLower41Comma": AccidentalCombiningLower41Comma,
	"accidentalCombiningLower43Comma": AccidentalCombiningLower43Comma,
	"accidentalCombiningLower47Quartertone": AccidentalCombiningLower47Quartertone,
	"accidentalCombiningLower53LimitComma": AccidentalCombiningLower53LimitComma,
	"accidentalCombiningOpenCurlyBrace": AccidentalCombiningOpenCurlyBrace,
	"accidentalCombiningRaise17Schisma": AccidentalCombiningRaise17Schisma,
	"accidentalCombiningRaise19Schisma": AccidentalCombiningRaise19Schisma,
	"accidentalCombiningRaise23Limit29LimitComma": AccidentalCombiningRaise23Limit29LimitComma,
	"accidentalCombiningRaise29LimitComma": AccidentalCombiningRaise29LimitComma,
	"accidentalCombiningRaise31Schisma": AccidentalCombiningRaise31Schisma,
	"accidentalCombiningRaise37Quartertone": AccidentalCombiningRaise37Quartertone,
	"accidentalCombiningRaise41Comma": AccidentalCombiningRaise41Comma,
	"accidentalCombiningRaise43Comma": AccidentalCombiningRaise43Comma,
	"accidentalCombiningRaise47Quartertone": AccidentalCombiningRaise47Quartertone,
	"accidentalCombiningRaise53LimitComma": AccidentalCombiningRaise53LimitComma,
	"accidentalCommaSlashDown": AccidentalCommaSlashDown,
	"accidentalCommaSlashUp": AccidentalCommaSlashUp,
	"accidentalDoubleFlat": AccidentalDoubleFlat,
	"accidentalDoubleFlatArabic": AccidentalDoubleFlatArabic,
	"accidentalDoubleFlatEqualTempered": AccidentalDoubleFlatEqualTempered,
	"accidentalDoubleFlatOneArrowDown": AccidentalDoubleFlatOneArrowDown,
	"accidentalDoubleFlatOneArrowUp": AccidentalDoubleFlatOneArrowUp,
	"accidentalDoubleFlatReversed": AccidentalDoubleFlatReversed,
	"accidentalDoubleFlatThreeArrowsDown": AccidentalDoubleFlatThreeArrowsDown,
	"accidentalDoubleFlatThreeArrowsUp": AccidentalDoubleFlatThreeArrowsUp,
	"accidentalDoubleFlatTurned": AccidentalDoubleFlatTurned,
	"accidentalDoubleFlatTwoArrowsDown": AccidentalDoubleFlatTwoArrowsDown,
	"accidentalDoubleFlatTwoArrowsUp": AccidentalDoubleFlatTwoArrowsUp,
	"accidentalDoubleSharp": AccidentalDoubleSharp,
	"accidentalDoubleSharpArabic": AccidentalDoubleSharpArabic,
	"accidentalDoubleSharpEqualTempered": AccidentalDoubleSharpEqualTempered,
	"accidentalDoubleSharpOneArrowDown": AccidentalDoubleSharpOneArrowDown,
	"accidentalDoubleSharpOneArrowUp": AccidentalDoubleSharpOneArrowUp,
	"accidentalDoubleSharpThreeArrowsDown": AccidentalDoubleSharpThreeArrowsDown,
	"accidentalDoubleSharpThreeArrowsUp": AccidentalDoubleSharpThreeArrowsUp,
	"accidentalDoubleSharpTwoArrowsDown": AccidentalDoubleSharpTwoArrowsDown,
	"accidentalDoubleSharpTwoArrowsUp": AccidentalDoubleSharpTwoArrowsUp,
	"accidentalEnharmonicAlmostEqualTo": AccidentalEnharmonicAlmostEqualTo,
	"accidentalEnharmonicEquals": AccidentalEnharmonicEquals,
	"accidentalEnharmonicTilde": AccidentalEnharmonicTilde,
	"accidentalFilledReversedFlatAndFlat": AccidentalFilledReversedFlatAndFlat,
	"accidentalFilledReversedFlatAndFlatArrowDown": AccidentalFilledReversedFlatAndFlatArrowDown,
	"accidentalFilledReversedFlatAndFlatArrowUp": AccidentalFilledReversedFlatAndFlatArrowUp,
	"accidentalFilledReversedFlatArrowDown": AccidentalFilledReversedFlatArrowDown,
	"accidentalFilledReversedFlatArrowUp": AccidentalFilledReversedFlatArrowUp,
	"accidentalFiveQuarterTonesFlatArrowDown": AccidentalFiveQuarterTonesFlatArrowDown,
	"accidentalFiveQuarterTonesSharpArrowUp": AccidentalFiveQuarterTonesSharpArrowUp,
	"accidentalFlat": AccidentalFlat,
	"accidentalFlatArabic": AccidentalFlatArabic,
	"accidentalFlatEqualTempered": AccidentalFlatEqualTempered,
	"accidentalFlatLoweredStockhausen": AccidentalFlatLoweredStockhausen,
	"accidentalFlatOneArrowDown": AccidentalFlatOneArrowDown,
	"accidentalFlatOneArrowUp": AccidentalFlatOneArrowUp,
	"accidentalFlatRaisedStockhausen": AccidentalFlatRaisedStockhausen,
	"accidentalFlatRepeatedLineStockhausen": AccidentalFlatRepeatedLineStockhausen,
	"accidentalFlatRepeatedSpaceStockhausen": AccidentalFlatRepeatedSpaceStockhausen,
	"accidentalFlatThreeArrowsDown": AccidentalFlatThreeArrowsDown,
	"accidentalFlatThreeArrowsUp": AccidentalFlatThreeArrowsUp,
	"accidentalFlatTurned": AccidentalFlatTurned,
	"accidentalFlatTwoArrowsDown": AccidentalFlatTwoArrowsDown,
	"accidentalFlatTwoArrowsUp": AccidentalFlatTwoArrowsUp,
	"accidentalHabaFlatQuarterToneHigher": AccidentalHabaFlatQuarterToneHigher,
	"accidentalHabaFlatThreeQuarterTonesLower": AccidentalHabaFlatThreeQuarterTonesLower,
	"accidentalHabaQuarterToneHigher": AccidentalHabaQuarterToneHigher,
	"accidentalHabaQuarterToneLower": AccidentalHabaQuarterToneLower,
	"accidentalHabaSharpQuarterToneLower": AccidentalHabaSharpQuarterToneLower,
	"accidentalHabaSharpThreeQuarterTonesHigher": AccidentalHabaSharpThreeQuarterTonesHigher,
	"accidentalHalfSharpArrowDown": AccidentalHalfSharpArrowDown,
	"accidentalHalfSharpArrowUp": AccidentalHalfSharpArrowUp,
	"accidentalJohnston13": AccidentalJohnston13,
	"accidentalJohnston31": AccidentalJohnston31,
	"accidentalJohnstonDown": AccidentalJohnstonDown,
	"accidentalJohnstonEl": AccidentalJohnstonEl,
	"accidentalJohnstonMinus": AccidentalJohnstonMinus,
	"accidentalJohnstonPlus": AccidentalJohnstonPlus,
	"accidentalJohnstonSeven": AccidentalJohnstonSeven,
	"accidentalJohnstonUp": AccidentalJohnstonUp,
	"accidentalKomaFlat": AccidentalKomaFlat,
	"accidentalKomaSharp": AccidentalKomaSharp,
	"accidentalKoron": AccidentalKoron,
	"accidentalKucukMucennebFlat": AccidentalKucukMucennebFlat,
	"accidentalKucukMucennebSharp": AccidentalKucukMucennebSharp,
	"accidentalLargeDoubleSharp": AccidentalLargeDoubleSharp,
	"accidentalLowerOneSeptimalComma": AccidentalLowerOneSeptimalComma,
	"accidentalLowerOneTridecimalQuartertone": AccidentalLowerOneTridecimalQuartertone,
	"accidentalLowerOneUndecimalQuartertone": AccidentalLowerOneUndecimalQuartertone,
	"accidentalLowerTwoSeptimalCommas": AccidentalLowerTwoSeptimalCommas,
	"accidentalLoweredStockhausen": AccidentalLoweredStockhausen,
	"accidentalNarrowReversedFlat": AccidentalNarrowReversedFlat,
	"accidentalNarrowReversedFlatAndFlat": AccidentalNarrowReversedFlatAndFlat,
	"accidentalNatural": AccidentalNatural,
	"accidentalNaturalArabic": AccidentalNaturalArabic,
	"accidentalNaturalEqualTempered": AccidentalNaturalEqualTempered,
	"accidentalNaturalFlat": AccidentalNaturalFlat,
	"accidentalNaturalLoweredStockhausen": AccidentalNaturalLoweredStockhausen,
	"accidentalNaturalOneArrowDown": AccidentalNaturalOneArrowDown,
	"accidentalNaturalOneArrowUp": AccidentalNaturalOneArrowUp,
	"accidentalNaturalRaisedStockhausen": AccidentalNaturalRaisedStockhausen,
	"accidentalNaturalReversed": AccidentalNaturalReversed,
	"accidentalNaturalSharp": AccidentalNaturalSharp,
	"accidentalNaturalThreeArrowsDown": AccidentalNaturalThreeArrowsDown,
	"accidentalNaturalThreeArrowsUp": AccidentalNaturalThreeArrowsUp,
	"accidentalNaturalTwoArrowsDown": AccidentalNaturalTwoArrowsDown,
	"accidentalNaturalTwoArrowsUp": AccidentalNaturalTwoArrowsUp,
	"accidentalOneAndAHalfSharpsArrowDown": AccidentalOneAndAHalfSharpsArrowDown,
	"accidentalOneAndAHalfSharpsArrowUp": AccidentalOneAndAHalfSharpsArrowUp,
	"accidentalOneQuarterToneFlatFerneyhough": AccidentalOneQuarterToneFlatFerneyhough,
	"accidentalOneQuarterToneFlatStockhausen": AccidentalOneQuarterToneFlatStockhausen,
	"accidentalOneQuarterToneSharpFerneyhough": AccidentalOneQuarterToneSharpFerneyhough,
	"accidentalOneQuarterToneSharpStockhausen": AccidentalOneQuarterToneSharpStockhausen,
	"accidentalOneThirdToneFlatFerneyhough": AccidentalOneThirdToneFlatFerneyhough,
	"accidentalOneThirdToneSharpFerneyhough": AccidentalOneThirdToneSharpFerneyhough,
	"accidentalParensLeft": AccidentalParensLeft,
	"accidentalParensRight": AccidentalParensRight,
	"accidentalQuarterFlatEqualTempered": AccidentalQuarterFlatEqualTempered,
	"accidentalQuarterSharpEqualTempered": AccidentalQuarterSharpEqualTempered,
	"accidentalQuarterToneFlat4": AccidentalQuarterToneFlat4,
	"accidentalQuarterToneFlatArabic": AccidentalQuarterToneFlatArabic,
	"accidentalQuarterToneFlatArrowUp": AccidentalQuarterToneFlatArrowUp,
	"accidentalQuarterToneFlatFilledReversed": AccidentalQuarterToneFlatFilledReversed,
	"accidentalQuarterToneFlatNaturalArrowDown": AccidentalQuarterToneFlatNaturalArrowDown,
	"accidentalQuarterToneFlatPenderecki": AccidentalQuarterToneFlatPenderecki,
	"accidentalQuarterToneFlatStein": AccidentalQuarterToneFlatStein,
	"accidentalQuarterToneFlatVanBlankenburg": AccidentalQuarterToneFlatVanBlankenburg,
	"accidentalQuarterToneSharp4": AccidentalQuarterToneSharp4,
	"accidentalQuarterToneSharpArabic": AccidentalQuarterToneSharpArabic,
	"accidentalQuarterToneSharpArrowDown": AccidentalQuarterToneSharpArrowDown,
	"accidentalQuarterToneSharpBusotti": AccidentalQuarterToneSharpBusotti,
	"accidentalQuarterToneSharpNaturalArrowUp": AccidentalQuarterToneSharpNaturalArrowUp,
	"accidentalQuarterToneSharpStein": AccidentalQuarterToneSharpStein,
	"accidentalQuarterToneSharpWiggle": AccidentalQuarterToneSharpWiggle,
	"accidentalRaiseOneSeptimalComma": AccidentalRaiseOneSeptimalComma,
	"accidentalRaiseOneTridecimalQuartertone": AccidentalRaiseOneTridecimalQuartertone,
	"accidentalRaiseOneUndecimalQuartertone": AccidentalRaiseOneUndecimalQuartertone,
	"accidentalRaiseTwoSeptimalCommas": AccidentalRaiseTwoSeptimalCommas,
	"accidentalRaisedStockhausen": AccidentalRaisedStockhausen,
	"accidentalReversedFlatAndFlatArrowDown": AccidentalReversedFlatAndFlatArrowDown,
	"accidentalReversedFlatAndFlatArrowUp": AccidentalReversedFlatAndFlatArrowUp,
	"accidentalReversedFlatArrowDown": AccidentalReversedFlatArrowDown,
	"accidentalReversedFlatArrowUp": AccidentalReversedFlatArrowUp,
	"accidentalSharp": AccidentalSharp,
	"accidentalSharpArabic": AccidentalSharpArabic,
	"accidentalSharpEqualTempered": AccidentalSharpEqualTempered,
	"accidentalSharpLoweredStockhausen": AccidentalSharpLoweredStockhausen,
	"accidentalSharpOneArrowDown": AccidentalSharpOneArrowDown,
	"accidentalSharpOneArrowUp": AccidentalSharpOneArrowUp,
	"accidentalSharpOneHorizontalStroke": AccidentalSharpOneHorizontalStroke,
	"accidentalSharpRaisedStockhausen": AccidentalSharpRaisedStockhausen,
	"accidentalSharpRepeatedLineStockhausen": AccidentalSharpRepeatedLineStockhausen,
	"accidentalSharpRepeatedSpaceStockhausen": AccidentalSharpRepeatedSpaceStockhausen,
	"accidentalSharpReversed": AccidentalSharpReversed,
	"accidentalSharpSharp": AccidentalSharpSharp,
	"accidentalSharpThreeArrowsDown": AccidentalSharpThreeArrowsDown,
	"accidentalSharpThreeArrowsUp": AccidentalSharpThreeArrowsUp,
	"accidentalSharpTwoArrowsDown": AccidentalSharpTwoArrowsDown,
	"accidentalSharpTwoArrowsUp": AccidentalSharpTwoArrowsUp,
	"accidentalSims12Down": AccidentalSims12Down,
	"accidentalSims12Up": AccidentalSims12Up,
	"accidentalSims4Down": AccidentalSims4Down,
	"accidentalSims4Up": AccidentalSims4Up,
	"accidentalSims6Down": AccidentalSims6Down,
	"accidentalSims6Up": AccidentalSims6Up,
	"accidentalSori": AccidentalSori,
	"accidentalTavenerFlat": AccidentalTavenerFlat,
	"accidentalTavenerSharp": AccidentalTavenerSharp,
	"accidentalThreeQuarterTonesFlatArabic": AccidentalThreeQuarterTonesFlatArabic,
	"accidentalThreeQuarterTonesFlatArrowDown": AccidentalThreeQuarterTonesFlatArrowDown,
	"accidentalThreeQuarterTonesFlatArrowUp": AccidentalThreeQuarterTonesFlatArrowUp,
	"accidentalThreeQuarterTonesFlatCouper": AccidentalThreeQuarterTonesFlatCouper,
	"accidentalThreeQuarterTonesFlatGrisey": AccidentalThreeQuarterTonesFlatGrisey,
	"accidentalThreeQuarterTonesFlatTartini": AccidentalThreeQuarterTonesFlatTartini,
	"accidentalThreeQuarterTonesFlatZimmermann": AccidentalThreeQuarterTonesFlatZimmermann,
	"accidentalThreeQuarterTonesSharpArabic": AccidentalThreeQuarterTonesSharpArabic,
	"accidentalThreeQuarterTonesSharpArrowDown": AccidentalThreeQuarterTonesSharpArrowDown,
	"accidentalThreeQuarterTonesSharpArrowUp": AccidentalThreeQuarterTonesSharpArrowUp,
	"accidentalThreeQuarterTonesSharpBusotti": AccidentalThreeQuarterTonesSharpBusotti,
	"accidentalThreeQuarterTonesSharpStein": AccidentalThreeQuarterTonesSharpStein,
	"accidentalThreeQuarterTonesSharpStockhausen": AccidentalThreeQuarterTonesSharpStockhausen,
	"accidentalTripleFlat": AccidentalTripleFlat,
	"accidentalTripleSharp": AccidentalTripleSharp,
	"accidentalTwoThirdTonesFlatFerneyhough": AccidentalTwoThirdTonesFlatFerneyhough,
	"accidentalTwoThirdTonesSharpFerneyhough": AccidentalTwoThirdTonesSharpFerneyhough,
	"accidentalUpsAndDownsDown": AccidentalUpsAndDownsDown,
	"accidentalUpsAndDownsLess": AccidentalUpsAndDownsLess,
	"accidentalUpsAndDownsMore": AccidentalUpsAndDownsMore,
	"accidentalUpsAndDownsUp": AccidentalUpsAndDownsUp,
	"accidentalWilsonMinus": AccidentalWilsonMinus,
	"accidentalWilsonPlus": AccidentalWilsonPlus,
	"accidentalWyschnegradsky10TwelfthsFlat": AccidentalWyschnegradsky10TwelfthsFlat,
	"accidentalWyschnegradsky10TwelfthsSharp": AccidentalWyschnegradsky10TwelfthsSharp,
	"accidentalWyschnegradsky11TwelfthsFlat": AccidentalWyschnegradsky11TwelfthsFlat,
	"accidentalWyschnegradsky11TwelfthsSharp": AccidentalWyschnegradsky11TwelfthsSharp,
	"accidentalWyschnegradsky1TwelfthsFlat": AccidentalWyschnegradsky1TwelfthsFlat,
	"accidentalWyschnegradsky1TwelfthsSharp": AccidentalWyschnegradsky1TwelfthsSharp,
	"accidentalWyschnegradsky2TwelfthsFlat": AccidentalWyschnegradsky2TwelfthsFlat,
	"accidentalWyschnegradsky2TwelfthsSharp": AccidentalWyschnegradsky2TwelfthsSharp,
	"accidentalWyschnegradsky3TwelfthsFlat": AccidentalWyschnegradsky3TwelfthsFlat,
	"accidentalWyschnegradsky3TwelfthsSharp": AccidentalWyschnegradsky3TwelfthsSharp,
	"accidentalWyschnegradsky4TwelfthsFlat": AccidentalWyschnegradsky4TwelfthsFlat,
	"accidentalWyschnegradsky4TwelfthsSharp": AccidentalWyschnegradsky4TwelfthsSharp,
	"accidentalWyschnegradsky5TwelfthsFlat": AccidentalWyschnegradsky5TwelfthsFlat,
	"accidentalWyschnegradsky5TwelfthsSharp": AccidentalWyschnegradsky5TwelfthsSharp,
	"accidentalWyschnegradsky6TwelfthsFlat": AccidentalWyschnegradsky6TwelfthsFlat,
	"accidentalWyschnegradsky6TwelfthsSharp": AccidentalWyschnegradsky6TwelfthsSharp,
	"accidentalWyschnegradsky7TwelfthsFlat": AccidentalWyschnegradsky7TwelfthsFlat,
	"accidentalWyschnegradsky7TwelfthsSharp": AccidentalWyschnegradsky7TwelfthsSharp,
	"accidentalWyschnegradsky8TwelfthsFlat": AccidentalWyschnegradsky8TwelfthsFlat,
	"accidentalWyschnegradsky8TwelfthsSharp": AccidentalWyschnegradsky8TwelfthsSharp,
	"accidentalWyschnegradsky9TwelfthsFlat": AccidentalWyschnegradsky9TwelfthsFlat,
	"accidentalWyschnegradsky9TwelfthsSharp": AccidentalWyschnegradsky9TwelfthsSharp,
	"accidentalXenakisOneThirdToneSharp": AccidentalXenakisOneThirdToneSharp,
	"accidentalXenakisTwoThirdTonesSharp": AccidentalXenakisTwoThirdTonesSharp,
	"analyticsChoralmelodie": AnalyticsChoralmelodie,
	"analyticsEndStimme": AnalyticsEndStimme,
	"analyticsHauptrhythmus": AnalyticsHauptrhythmus,
	"analyticsHauptstimme": AnalyticsHauptstimme,
	"analyticsInversion1": AnalyticsInversion1,
	"analyticsNebenstimme": AnalyticsNebenstimme,
	"analyticsStartStimme": AnalyticsStartStimme,
	"analyticsTheme": AnalyticsTheme,
	"analyticsTheme1": AnalyticsTheme1,
	"analyticsThemeInversion": AnalyticsThemeInversion,
	"analyticsThemeRetrograde": AnalyticsThemeRetrograde,
	"analyticsThemeRetrogradeInversion": AnalyticsThemeRetrogradeInversion,
	"arpeggiato": Arpeggiato,
	"arpeggiatoDown": ArpeggiatoDown,
	"arpeggiatoUp": ArpeggiatoUp,
	"arrowBlackDown": ArrowBlackDown,
	"arrowBlackDownLeft": ArrowBlackDownLeft,
	"arrowBlackDownRight": ArrowBlackDownRight,
	"arrowBlackLeft": ArrowBlackLeft,
	"arrowBlackRight": ArrowBlackRight,
	"arrowBlackUp": ArrowBlackUp,
	"arrowBlackUpLeft": ArrowBlackUpLeft,
	"arrowBlackUpRight": ArrowBlackUpRight,
	"arrowOpenDown": ArrowOpenDown,
	"arrowOpenDownLeft": ArrowOpenDownLeft,
	"arrowOpenDownRight": ArrowOpenDownRight,
	"arrowOpenLeft": ArrowOpenLeft,
	"arrowOpenRight": ArrowOpenRight,
	"arrowOpenUp": ArrowOpenUp,
	"arrowOpenUpLeft": ArrowOpenUpLeft,
	"arrowOpenUpRight": ArrowOpenUpRight,
	"arrowWhiteDown": ArrowWhiteDown,
	"arrowWhiteDownLeft": ArrowWhiteDownLeft,
	"arrowWhiteDownRight": ArrowWhiteDownRight,
	"arrowWhiteLeft": ArrowWhiteLeft,
	"arrowWhiteRight": ArrowWhiteRight,
	"arrowWhiteUp": ArrowWhiteUp,
	"arrowWhiteUpLeft": ArrowWhiteUpLeft,
	"arrowWhiteUpRight": ArrowWhiteUpRight,
	"arrowheadBlackDown": ArrowheadBlackDown,
	"arrowheadBlackDownLeft": ArrowheadBlackDownLeft,
	"arrowheadBlackDownRight": ArrowheadBlackDownRight,
	"arrowheadBlackLeft": ArrowheadBlackLeft,
	"arrowheadBlackRight": ArrowheadBlackRight,
	"arrowheadBlackUp": ArrowheadBlackUp,
	"arrowheadBlackUpLeft": ArrowheadBlackUpLeft,
	"arrowheadBlackUpRight": ArrowheadBlackUpRight,
	"arrowheadOpenDown": ArrowheadOpenDown,
	"arrowheadOpenDownLeft": ArrowheadOpenDownLeft,
	"arrowheadOpenDownRight": ArrowheadOpenDownRight,
	"arrowheadOpenLeft": ArrowheadOpenLeft,
	"arrowheadOpenRight": ArrowheadOpenRight,
	"arrowheadOpenUp": ArrowheadOpenUp,
	"arrowheadOpenUpLeft": ArrowheadOpenUpLeft,
	"arrowheadOpenUpRight": ArrowheadOpenUpRight,
	"arrowheadWhiteDown": ArrowheadWhiteDown,
	"arrowheadWhiteDownLeft": ArrowheadWhiteDownLeft,
	"arrowheadWhiteDownRight": ArrowheadWhiteDownRight,
	"arrowheadWhiteLeft": ArrowheadWhiteLeft,
	"arrowheadWhiteRight": ArrowheadWhiteRight,
	"arrowheadWhiteUp": ArrowheadWhiteUp,
	"arrowheadWhiteUpLeft": ArrowheadWhiteUpLeft,
	"arrowheadWhiteUpRight": ArrowheadWhiteUpRight,
	"articAccentAbove": ArticAccentAbove,
	"articAccentBelow": ArticAccentBelow,
	"articAccentStaccatoAbove": ArticAccentStaccatoAbove,
	"articAccentStaccatoBelow": ArticAccentStaccatoBelow,
	"articLaissezVibrerAbove": ArticLaissezVibrerAbove,
	"articLaissezVibrerBelow": ArticLaissezVibrerBelow,
	"articMarcatoAbove": ArticMarcatoAbove,
	"articMarcatoBelow": ArticMarcatoBelow,
	"articMarcatoStaccatoAbove": ArticMarcatoStaccatoAbove,
	"articMarcatoStaccatoBelow": ArticMarcatoStaccatoBelow,
	"articMarcatoTenutoAbove": ArticMarcatoTenutoAbove,
	"articMarcatoTenutoBelow": ArticMarcatoTenutoBelow,
	"articSoftAccentAbove": ArticSoftAccentAbove,
	"articSoftAccentBelow": ArticSoftAccentBelow,
	"articSoftAccentStaccatoAbove": ArticSoftAccentStaccatoAbove,
	"articSoftAccentStaccatoBelow": ArticSoftAccentStaccatoBelow,
	"articSoftAccentTenutoAbove": ArticSoftAccentTenutoAbove,
	"articSoftAccentTenutoBelow": ArticSoftAccentTenutoBelow,
	"articSoftAccentTenutoStaccatoAbove": ArticSoftAccentTenutoStaccatoAbove,
	"articSoftAccentTenutoStaccatoBelow": ArticSoftAccentTenutoStaccatoBelow,
	"articStaccatissimoAbove": ArticStaccatissimoAbove,
	"articStaccatissimoBelow": ArticStaccatissimoBelow,
	"articStaccatissimoStrokeAbove": ArticStaccatissimoStrokeAbove,
	"articStaccatissimoStrokeBelow": ArticStaccatissimoStrokeBelow,
	"articStaccatissimoWedgeAbove": ArticStaccatissimoWedgeAbove,
	"articStaccatissimoWedgeBelow": ArticStaccatissimoWedgeBelow,
	"articStaccatoAbove": ArticStaccatoAbove,
	"articStaccatoBelow": ArticStaccatoBelow,
	"articStressAbove": ArticStressAbove,
	"articStressBelow": ArticStressBelow,
	"articTenutoAbove": ArticTenutoAbove,
	"articTenutoAccentAbove": ArticTenutoAccentAbove,
	"articTenutoAccentBelow": ArticTenutoAccentBelow,
	"articTenutoBelow": ArticTenutoBelow,
	"articTenutoStaccatoAbove": ArticTenutoStaccatoAbove,
	"articTenutoStaccatoBelow": ArticTenutoStaccatoBelow,
	"articUnstressAbove": ArticUnstressAbove,
	"articUnstressBelow": ArticUnstressBelow,
	"augmentationDot": AugmentationDot,
	"barlineDashed": BarlineDashed,
	"barlineDotted": BarlineDotted,
	"barlineDouble": BarlineDouble,
	"barlineFinal": BarlineFinal,
	"barlineHeavy": BarlineHeavy,
	"barlineHeavyHeavy": BarlineHeavyHeavy,
	"barlineReverseFinal": BarlineReverseFinal,
	"barlineShort": BarlineShort,
	"barlineSingle": BarlineSingle,
	"barlineTick": BarlineTick,
	"beamAccelRit1": BeamAccelRit1,
	"beamAccelRit10": BeamAccelRit10,
	"beamAccelRit11": BeamAccelRit11,
	"beamAccelRit12": BeamAccelRit12,
	"beamAccelRit13": BeamAccelRit13,
	"beamAccelRit14": BeamAccelRit14,
	"beamAccelRit15": BeamAccelRit15,
	"beamAccelRit2": BeamAccelRit2,
	"beamAccelRit3": BeamAccelRit3,
	"beamAccelRit4": BeamAccelRit4,
	"beamAccelRit5": BeamAccelRit5,
	"beamAccelRit6": BeamAccelRit6,
	"beamAccelRit7": BeamAccelRit7,
	"beamAccelRit8": BeamAccelRit8,
	"beamAccelRit9": BeamAccelRit9,
	"beamAccelRitFinal": BeamAccelRitFinal,
	"brace": Brace,
	"bracket": Bracket,
	"bracketBottom": BracketBottom,
	"bracketTop": BracketTop,
	"brassBend": BrassBend,
	"brassDoitLong": BrassDoitLong,
	"brassDoitMedium": BrassDoitMedium,
	"brassDoitShort": BrassDoitShort,
	"brassFallLipLong": BrassFallLipLong,
	"brassFallLipMedium": BrassFallLipMedium,
	"brassFallLipShort": BrassFallLipShort,
	"brassFallRoughLong": BrassFallRoughLong,
	"brassFallRoughMedium": BrassFallRoughMedium,
	"brassFallRoughShort": BrassFallRoughShort,
	"brassFallSmoothLong": BrassFallSmoothLong,
	"brassFallSmoothMedium": BrassFallSmoothMedium,
	"brassFallSmoothShort": BrassFallSmoothShort,
	"brassFlip": BrassFlip,
	"brassHarmonMuteClosed": BrassHarmonMuteClosed,
	"brassHarmonMuteStemHalfLeft": BrassHarmonMuteStemHalfLeft,
	"brassHarmonMuteStemHalfRight": BrassHarmonMuteStemHalfRight,
	"brassHarmonMuteStemOpen": BrassHarmonMuteStemOpen,
	"brassJazzTurn": BrassJazzTurn,
	"brassLiftLong": BrassLiftLong,
	"brassLiftMedium": BrassLiftMedium,
	"brassLiftShort": BrassLiftShort,
	"brassLiftSmoothLong": BrassLiftSmoothLong,
	"brassLiftSmoothMedium": BrassLiftSmoothMedium,
	"brassLiftSmoothShort": BrassLiftSmoothShort,
	"brassMuteClosed": BrassMuteClosed,
	"brassMuteHalfClosed": BrassMuteHalfClosed,
	"brassMuteOpen": BrassMuteOpen,
	"brassPlop": BrassPlop,
	"brassScoop": BrassScoop,
	"brassSmear": BrassSmear,
	"brassValveTrill": BrassValveTrill,
	"breathMarkComma": BreathMarkComma,
	"breathMarkSalzedo": BreathMarkSalzedo,
	"breathMarkTick": BreathMarkTick,
	"breathMarkUpbow": BreathMarkUpbow,
	"bridgeClef": BridgeClef,
	"buzzRoll": BuzzRoll,
	"cClef": CClef,
	"cClef8vb": CClef8vb,
	"cClefArrowDown": CClefArrowDown,
	"cClefArrowUp": CClefArrowUp,
	"cClefChange": CClefChange,
	"cClefCombining": CClefCombining,
	"cClefReversed": CClefReversed,
	"cClefSquare": CClefSquare,
	"caesura": Caesura,
	"caesuraCurved": CaesuraCurved,
	"caesuraShort": CaesuraShort,
	"caesuraSingleStroke": CaesuraSingleStroke,
	"caesuraThick": CaesuraThick,
	"chantAccentusAbove": ChantAccentusAbove,
	"chantAccentusBelow": ChantAccentusBelow,
	"chantAuctumAsc": ChantAuctumAsc,
	"chantAuctumDesc": ChantAuctumDesc,
	"chantAugmentum": ChantAugmentum,
	"chantCaesura": ChantCaesura,
	"chantCclef": ChantCclef,
	"chantCirculusAbove": ChantCirculusAbove,
	"chantCirculusBelow": ChantCirculusBelow,
	"chantConnectingLineAsc2nd": ChantConnectingLineAsc2nd,
	"chantConnectingLineAsc3rd": ChantConnectingLineAsc3rd,
	"chantConnectingLineAsc4th": ChantConnectingLineAsc4th,
	"chantConnectingLineAsc5th": ChantConnectingLineAsc5th,
	"chantConnectingLineAsc6th": ChantConnectingLineAsc6th,
	"chantCustosStemDownPosHigh": ChantCustosStemDownPosHigh,
	"chantCustosStemDownPosHighest": ChantCustosStemDownPosHighest,
	"chantCustosStemDownPosMiddle": ChantCustosStemDownPosMiddle,
	"chantCustosStemUpPosLow": ChantCustosStemUpPosLow,
	"chantCustosStemUpPosLowest": ChantCustosStemUpPosLowest,
	"chantCustosStemUpPosMiddle": ChantCustosStemUpPosMiddle,
	"chantDeminutumLower": ChantDeminutumLower,
	"chantDeminutumUpper": ChantDeminutumUpper,
	"chantDivisioFinalis": ChantDivisioFinalis,
	"chantDivisioMaior": ChantDivisioMaior,
	"chantDivisioMaxima": ChantDivisioMaxima,
	"chantDivisioMinima": ChantDivisioMinima,
	"chantEntryLineAsc2nd": ChantEntryLineAsc2nd,
	"chantEntryLineAsc3rd": ChantEntryLineAsc3rd,
	"chantEntryLineAsc4th": ChantEntryLineAsc4th,
	"chantEntryLineAsc5th": ChantEntryLineAsc5th,
	"chantEntryLineAsc6th": ChantEntryLineAsc6th,
	"chantEpisema": ChantEpisema,
	"chantFclef": ChantFclef,
	"chantIctusAbove": ChantIctusAbove,
	"chantIctusBelow": ChantIctusBelow,
	"chantLigaturaDesc2nd": ChantLigaturaDesc2nd,
	"chantLigaturaDesc3rd": ChantLigaturaDesc3rd,
	"chantLigaturaDesc4th": ChantLigaturaDesc4th,
	"chantLigaturaDesc5th": ChantLigaturaDesc5th,
	"chantOriscusAscending": ChantOriscusAscending,
	"chantOriscusDescending": ChantOriscusDescending,
	"chantOriscusLiquescens": ChantOriscusLiquescens,
	"chantPodatusLower": ChantPodatusLower,
	"chantPodatusUpper": ChantPodatusUpper,
	"chantPunctum": ChantPunctum,
	"chantPunctumCavum": ChantPunctumCavum,
	"chantPunctumDeminutum": ChantPunctumDeminutum,
	"chantPunctumInclinatum": ChantPunctumInclinatum,
	"chantPunctumInclinatumAuctum": ChantPunctumInclinatumAuctum,
	"chantPunctumInclinatumDeminutum": ChantPunctumInclinatumDeminutum,
	"chantPunctumLinea": ChantPunctumLinea,
	"chantPunctumLineaCavum": ChantPunctumLineaCavum,
	"chantPunctumVirga": ChantPunctumVirga,
	"chantPunctumVirgaReversed": ChantPunctumVirgaReversed,
	"chantQuilisma": ChantQuilisma,
	"chantSemicirculusAbove": ChantSemicirculusAbove,
	"chantSemicirculusBelow": ChantSemicirculusBelow,
	"chantStaff": ChantStaff,
	"chantStaffNarrow": ChantStaffNarrow,
	"chantStaffWide": ChantStaffWide,
	"chantStrophicus": ChantStrophicus,
	"chantStrophicusAuctus": ChantStrophicusAuctus,
	"chantStrophicusLiquescens2nd": ChantStrophicusLiquescens2nd,
	"chantStrophicusLiquescens3rd": ChantStrophicusLiquescens3rd,
	"chantStrophicusLiquescens4th": ChantStrophicusLiquescens4th,
	"chantStrophicusLiquescens5th": ChantStrophicusLiquescens5th,
	"chantVirgula": ChantVirgula,
	"clef15": Clef15,
	"clef8": Clef8,
	"clefChangeCombining": ClefChangeCombining,
	"coda": Coda,
	"codaSquare": CodaSquare,
	"conductorBeat2Compound": ConductorBeat2Compound,
	"conductorBeat2Simple": ConductorBeat2Simple,
	"conductorBeat3Compound": ConductorBeat3Compound,
	"conductorBeat3Simple": ConductorBeat3Simple,
	"conductorBeat4Compound": ConductorBeat4Compound,
	"conductorBeat4Simple": ConductorBeat4Simple,
	"conductorLeftBeat": ConductorLeftBeat,
	"conductorRightBeat": ConductorRightBeat,
	"conductorStrongBeat": ConductorStrongBeat,
	"conductorUnconducted": ConductorUnconducted,
	"conductorWeakBeat": ConductorWeakBeat,
	"controlBeginBeam": ControlBeginBeam,
	"controlBeginPhrase": ControlBeginPhrase,
	"controlBeginSlur": ControlBeginSlur,
	"controlBeginTie": ControlBeginTie,
	"controlEndBeam": ControlEndBeam,
	"controlEndPhrase": ControlEndPhrase,
	"controlEndSlur": ControlEndSlur,
	"controlEndTie": ControlEndTie,
	"csymAccidentalDoubleFlat": CsymAccidentalDoubleFlat,
	"csymAccidentalDoubleSharp": CsymAccidentalDoubleSharp,
	"csymAccidentalFlat": CsymAccidentalFlat,
	"csymAccidentalNatural": CsymAccidentalNatural,
	"csymAccidentalSharp": CsymAccidentalSharp,
	"csymAccidentalTripleFlat": CsymAccidentalTripleFlat,
	"csymAccidentalTripleSharp": CsymAccidentalTripleSharp,
	"csymAlteredBassSlash": CsymAlteredBassSlash,
	"csymAugmented": CsymAugmented,
	"csymBracketLeftTall": CsymBracketLeftTall,
	"csymBracketRightTall": CsymBracketRightTall,
	"csymDiagonalArrangementSlash": CsymDiagonalArrangementSlash,
	"csymDiminished": CsymDiminished,
	"csymHalfDiminished": CsymHalfDiminished,
	"csymMajorSeventh": CsymMajorSeventh,
	"csymMinor": CsymMinor,
	"csymParensLeftTall": CsymParensLeftTall,
	"csymParensLeftVeryTall": CsymParensLeftVeryTall,
	"csymParensRightTall": CsymParensRightTall,
	"csymParensRightVeryTall": CsymParensRightVeryTall,
	"curlewSign": CurlewSign,
	"daCapo": DaCapo,
	"dalSegno": DalSegno,
	"daseianExcellentes1": DaseianExcellentes1,
	"daseianExcellentes2": DaseianExcellentes2,
	"daseianExcellentes3": DaseianExcellentes3,
	"daseianExcellentes4": DaseianExcellentes4,
	"daseianFinales1": DaseianFinales1,
	"daseianFinales2": DaseianFinales2,
	"daseianFinales3": DaseianFinales3,
	"daseianFinales4": DaseianFinales4,
	"daseianGraves1": DaseianGraves1,
	"daseianGraves2": DaseianGraves2,
	"daseianGraves3": DaseianGraves3,
	"daseianGraves4": DaseianGraves4,
	"daseianResidua1": DaseianResidua1,
	"daseianResidua2": DaseianResidua2,
	"daseianSuperiores1": DaseianSuperiores1,
	"daseianSuperiores2": DaseianSuperiores2,
	"daseianSuperiores3": DaseianSuperiores3,
	"daseianSuperiores4": DaseianSuperiores4,
	"doubleLateralRollStevens": DoubleLateralRollStevens,
	"doubleTongueAbove": DoubleTongueAbove,
	"doubleTongueBelow": DoubleTongueBelow,
	"dynamicCombinedSeparatorColon": DynamicCombinedSeparatorColon,
	"dynamicCombinedSeparatorHyphen": DynamicCombinedSeparatorHyphen,
	"dynamicCombinedSeparatorSlash": DynamicCombinedSeparatorSlash,
	"dynamicCombinedSeparatorSpace": DynamicCombinedSeparatorSpace,
	"dynamicCrescendoHairpin": DynamicCrescendoHairpin,
	"dynamicDiminuendoHairpin": DynamicDiminuendoHairpin,
	"dynamicFF": DynamicFF,
	"dynamicFFF": DynamicFFF,
	"dynamicFFFF": DynamicFFFF,
	"dynamicFFFFF": DynamicFFFFF,
	"dynamicFFFFFF": DynamicFFFFFF,
	"dynamicForte": DynamicForte,
	"dynamicFortePiano": DynamicFortePiano,
	"dynamicForzando": DynamicForzando,
	"dynamicHairpinBracketLeft": DynamicHairpinBracketLeft,
	"dynamicHairpinBracketRight": DynamicHairpinBracketRight,
	"dynamicHairpinParenthesisLeft": DynamicHairpinParenthesisLeft,
	"dynamicHairpinParenthesisRight": DynamicHairpinParenthesisRight,
	"dynamicMF": DynamicMF,
	"dynamicMP": DynamicMP,
	"dynamicMessaDiVoce": DynamicMessaDiVoce,
	"dynamicMezzo": DynamicMezzo,
	"dynamicNiente": DynamicNiente,
	"dynamicNienteForHairpin": DynamicNienteForHairpin,
	"dynamicPF": DynamicPF,
	"dynamicPP": DynamicPP,
	"dynamicPPP": DynamicPPP,
	"dynamicPPPP": DynamicPPPP,
	"dynamicPPPPP": DynamicPPPPP,
	"dynamicPPPPPP": DynamicPPPPPP,
	"dynamicPiano": DynamicPiano,
	"dynamicRinforzando": DynamicRinforzando,
	"dynamicRinforzando1": DynamicRinforzando1,
	"dynamicRinforzando2": DynamicRinforzando2,
	"dynamicSforzando": DynamicSforzando,
	"dynamicSforzando1": DynamicSforzando1,
	"dynamicSforzandoPianissimo": DynamicSforzandoPianissimo,
	"dynamicSforzandoPiano": DynamicSforzandoPiano,
	"dynamicSforzato": DynamicSforzato,
	"dynamicSforzatoFF": DynamicSforzatoFF,
	"dynamicSforzatoPiano": DynamicSforzatoPiano,
	"dynamicZ": DynamicZ,
	"elecAudioChannelsEight": ElecAudioChannelsEight,
	"elecAudioChannelsFive": ElecAudioChannelsFive,
	"elecAudioChannelsFour": ElecAudioChannelsFour,
	"elecAudioChannelsOne": ElecAudioChannelsOne,
	"elecAudioChannelsSeven": ElecAudioChannelsSeven,
	"elecAudioChannelsSix": ElecAudioChannelsSix,
	"elecAudioChannelsThreeFrontal": ElecAudioChannelsThreeFrontal,
	"elecAudioChannelsThreeSurround": ElecAudioChannelsThreeSurround,
	"elecAudioChannelsTwo": ElecAudioChannelsTwo,
	"elecAudioIn": ElecAudioIn,
	"elecAudioMono": ElecAudioMono,
	"elecAudioOut": ElecAudioOut,
	"elecAudioStereo": ElecAudioStereo,
	"elecCamera": ElecCamera,
	"elecDataIn": ElecDataIn,
	"elecDataOut": ElecDataOut,
	"elecDisc": ElecDisc,
	"elecDownload": ElecDownload,
	"elecEject": ElecEject,
	"elecFastForward": ElecFastForward,
	"elecHeadphones": ElecHeadphones,
	"elecHeadset": ElecHeadset,
	"elecLineIn": ElecLineIn,
	"elecLineOut": ElecLineOut,
	"elecLoop": ElecLoop,
	"elecLoudspeaker": ElecLoudspeaker,
	"elecMIDIController0": ElecMIDIController0,
	"elecMIDIController100": ElecMIDIController100,
	"elecMIDIController20": ElecMIDIController20,
	"elecMIDIController40": ElecMIDIController40,
	"elecMIDIController60": ElecMIDIController60,
	"elecMIDIController80": ElecMIDIController80,
	"elecMIDIIn": ElecMIDIIn,
	"elecMIDIOut": ElecMIDIOut,
	"elecMicrophone": ElecMicrophone,
	"elecMicrophoneMute": ElecMicrophoneMute,
	"elecMicrophoneUnmute": ElecMicrophoneUnmute,
	"elecMixingConsole": ElecMixingConsole,
	"elecMonitor": ElecMonitor,
	"elecMute": ElecMute,
	"elecPause": ElecPause,
	"elecPlay": ElecPlay,
	"elecPowerOnOff": ElecPowerOnOff,
	"elecProjector": ElecProjector,
	"elecReplay": ElecReplay,
	"elecRewind": ElecRewind,
	"elecShuffle": ElecShuffle,
	"elecSkipBackwards": ElecSkipBackwards,
	"elecSkipForwards": ElecSkipForwards,
	"elecStop": ElecStop,
	"elecTape": ElecTape,
	"elecUSB": ElecUSB,
	"elecUnmute": ElecUnmute,
	"elecUpload": ElecUpload,
	"elecVideoCamera": ElecVideoCamera,
	"elecVideoIn": ElecVideoIn,
	"elecVideoOut": ElecVideoOut,
	"elecVolumeFader": ElecVolumeFader,
	"elecVolumeFaderThumb": ElecVolumeFaderThumb,
	"elecVolumeLevel0": ElecVolumeLevel0,
	"elecVolumeLevel100": ElecVolumeLevel100,
	"elecVolumeLevel20": ElecVolumeLevel20,
	"elecVolumeLevel40": ElecVolumeLevel40,
	"elecVolumeLevel60": ElecVolumeLevel60,
	"elecVolumeLevel80": ElecVolumeLevel80,
	"fClef": FClef,
	"fClef15ma": FClef15ma,
	"fClef15mb": FClef15mb,
	"fClef8va": FClef8va,
	"fClef8vb": FClef8vb,
	"fClefArrowDown": FClefArrowDown,
	"fClefArrowUp": FClefArrowUp,
	"fClefChange": FClefChange,
	"fClefReversed": FClefReversed,
	"fClefTurned": FClefTurned,
	"fermataAbove": FermataAbove,
	"fermataBelow": FermataBelow,
	"fermataLongAbove": FermataLongAbove,
	"fermataLongBelow": FermataLongBelow,
	"fermataLongHenzeAbove": FermataLongHenzeAbove,
	"fermataLongHenzeBelow": FermataLongHenzeBelow,
	"fermataShortAbove": FermataShortAbove,
	"fermataShortBelow": FermataShortBelow,
	"fermataShortHenzeAbove": FermataShortHenzeAbove,
	"fermataShortHenzeBelow": FermataShortHenzeBelow,
	"fermataVeryLongAbove": FermataVeryLongAbove,
	"fermataVeryLongBelow": FermataVeryLongBelow,
	"fermataVeryShortAbove": FermataVeryShortAbove,
	"fermataVeryShortBelow": FermataVeryShortBelow,
	"figbass0": Figbass0,
	"figbass1": Figbass1,
	"figbass2": Figbass2,
	"figbass2Raised": Figbass2Raised,
	"figbass3": Figbass3,
	"figbass4": Figbass4,
	"figbass4Raised": Figbass4Raised,
	"figbass5": Figbass5,
	"figbass5Raised1": Figbass5Raised1,
	"figbass5Raised2": Figbass5Raised2,
	"figbass5Raised3": Figbass5Raised3,
	"figbass6": Figbass6,
	"figbass6Raised": Figbass6Raised,
	"figbass6Raised2": Figbass6Raised2,
	"figbass7": Figbass7,
	"figbass7Diminished": Figbass7Diminished,
	"figbass7Raised1": Figbass7Raised1,
	"figbass7Raised2": Figbass7Raised2,
	"figbass8": Figbass8,
	"figbass9": Figbass9,
	"figbass9Raised": Figbass9Raised,
	"figbassBracketLeft": FigbassBracketLeft,
	"figbassBracketRight": FigbassBracketRight,
	"figbassCombiningLowering": FigbassCombiningLowering,
	"figbassCombiningRaising": FigbassCombiningRaising,
	"figbassDoubleFlat": FigbassDoubleFlat,
	"figbassDoubleSharp": FigbassDoubleSharp,
	"figbassFlat": FigbassFlat,
	"figbassNatural": FigbassNatural,
	"figbassParensLeft": FigbassParensLeft,
	"figbassParensRight": FigbassParensRight,
	"figbassPlus": FigbassPlus,
	"figbassSharp": FigbassSharp,
	"figbassTripleFlat": FigbassTripleFlat,
	"figbassTripleSharp": FigbassTripleSharp,
	"fingering0": Fingering0,
	"fingering0Italic": Fingering0Italic,
	"fingering1": Fingering1,
	"fingering1Italic": Fingering1Italic,
	"fingering2": Fingering2,
	"fingering2Italic": Fingering2Italic,
	"fingering3": Fingering3,
	"fingering3Italic": Fingering3Italic,
	"fingering4": Fingering4,
	"fingering4Italic": Fingering4Italic,
	"fingering5": Fingering5,
	"fingering5Italic": Fingering5Italic,
	"fingering6": Fingering6,
	"fingering6Italic": Fingering6Italic,
	"fingering7": Fingering7,
	"fingering7Italic": Fingering7Italic,
	"fingering8": Fingering8,
	"fingering8Italic": Fingering8Italic,
	"fingering9": Fingering9,
	"fingering9Italic": Fingering9Italic,
	"fingeringALower": FingeringALower,
	"fingeringCLower": FingeringCLower,
	"fingeringELower": FingeringELower,
	"fingeringILower": FingeringILower,
	"fingeringLeftBracket": FingeringLeftBracket,
	"fingeringLeftBracketItalic": FingeringLeftBracketItalic,
	"fingeringLeftParenthesis": FingeringLeftParenthesis,
	"fingeringLeftParenthesisItalic": FingeringLeftParenthesisItalic,
	"fingeringMLower": FingeringMLower,
	"fingeringMultipleNotes": FingeringMultipleNotes,
	"fingeringOLower": FingeringOLower,
	"fingeringPLower": FingeringPLower,
	"fingeringQLower": FingeringQLower,
	"fingeringRightBracket": FingeringRightBracket,
	"fingeringRightBracketItalic": FingeringRightBracketItalic,
	"fingeringRightParenthesis": FingeringRightParenthesis,
	"fingeringRightParenthesisItalic": FingeringRightParenthesisItalic,
	"fingeringSLower": FingeringSLower,
	"fingeringSeparatorMiddleDot": FingeringSeparatorMiddleDot,
	"fingeringSeparatorMiddleDotWhite": FingeringSeparatorMiddleDotWhite,
	"fingeringSeparatorSlash": FingeringSeparatorSlash,
	"fingeringSubstitutionAbove": FingeringSubstitutionAbove,
	"fingeringSubstitutionBelow": FingeringSubstitutionBelow,
	"fingeringSubstitutionDash": FingeringSubstitutionDash,
	"fingeringTLower": FingeringTLower,
	"fingeringTUpper": FingeringTUpper,
	"fingeringXLower": FingeringXLower,
	"flag1024thDown": Flag1024thDown,
	"flag1024thUp": Flag1024thUp,
	"flag128thDown": Flag128thDown,
	"flag128thUp": Flag128thUp,
	"flag16thDown": Flag16thDown,
	"flag16thUp": Flag16thUp,
	"flag256thDown": Flag256thDown,
	"flag256thUp": Flag256thUp,
	"flag32ndDown": Flag32ndDown,
	"flag32ndUp": Flag32ndUp,
	"flag512thDown": Flag512thDown,
	"flag512thUp": Flag512thUp,
	"flag64thDown": Flag64thDown,
	"flag64thUp": Flag64thUp,
	"flag8thDown": Flag8thDown,
	"flag8thUp": Flag8thUp,
	"flagInternalDown": FlagInternalDown,
	"flagInternalUp": FlagInternalUp,
	"fretboard3String": Fretboard3String,
	"fretboard3StringNut": Fretboard3StringNut,
	"fretboard4String": Fretboard4String,
	"fretboard4StringNut": Fretboard4StringNut,
	"fretboard5String": Fretboard5String,
	"fretboard5StringNut": Fretboard5StringNut,
	"fretboard6String": Fretboard6String,
	"fretboard6StringNut": Fretboard6StringNut,
	"fretboardFilledCircle": FretboardFilledCircle,
	"fretboardO": FretboardO,
	"fretboardX": FretboardX,
	"functionAngleLeft": FunctionAngleLeft,
	"functionAngleRight": FunctionAngleRight,
	"functionBracketLeft": FunctionBracketLeft,
	"functionBracketRight": FunctionBracketRight,
	"functionDD": FunctionDD,
	"functionDLower": FunctionDLower,
	"functionDUpper": FunctionDUpper,
	"functionEight": FunctionEight,
	"functionFUpper": FunctionFUpper,
	"functionFive": FunctionFive,
	"functionFour": FunctionFour,
	"functionGLower": FunctionGLower,
	"functionGUpper": FunctionGUpper,
	"functionGreaterThan": FunctionGreaterThan,
	"functionILower": FunctionILower,
	"functionIUpper": FunctionIUpper,
	"functionKLower": FunctionKLower,
	"functionKUpper": FunctionKUpper,
	"functionLLower": FunctionLLower,
	"functionLUpper": FunctionLUpper,
	"functionLessThan": FunctionLessThan,
	"functionMLower": FunctionMLower,
	"functionMUpper": FunctionMUpper,
	"functionMinus": FunctionMinus,
	"functionNLower": FunctionNLower,
	"functionNUpper": FunctionNUpper,
	"functionNUpperSuperscript": FunctionNUpperSuperscript,
	"functionNine": FunctionNine,
	"functionOne": FunctionOne,
	"functionPLower": FunctionPLower,
	"functionPUpper": FunctionPUpper,
	"functionParensLeft": FunctionParensLeft,
	"functionParensRight": FunctionParensRight,
	"functionPlus": FunctionPlus,
	"functionRLower": FunctionRLower,
	"functionRepetition1": FunctionRepetition1,
	"functionRepetition2": FunctionRepetition2,
	"functionRing": FunctionRing,
	"functionSLower": FunctionSLower,
	"functionSSLower": FunctionSSLower,
	"functionSSUpper": FunctionSSUpper,
	"functionSUpper": FunctionSUpper,
	"functionSeven": FunctionSeven,
	"functionSix": FunctionSix,
	"functionSlashedDD": FunctionSlashedDD,
	"functionTLower": FunctionTLower,
	"functionTUpper": FunctionTUpper,
	"functionThree": FunctionThree,
	"functionTwo": FunctionTwo,
	"functionVLower": FunctionVLower,
	"functionVUpper": FunctionVUpper,
	"functionZero": FunctionZero,
	"gClef": GClef,
	"gClef15ma": GClef15ma,
	"gClef15mb": GClef15mb,
	"gClef8va": GClef8va,
	"gClef8vb": GClef8vb,
	"gClef8vbCClef": GClef8vbCClef,
	"gClef8vbOld": GClef8vbOld,
	"gClef8vbParens": GClef8vbParens,
	"gClefArrowDown": GClefArrowDown,
	"gClefArrowUp": GClefArrowUp,
	"gClefChange": GClefChange,
	"gClefLigatedNumberAbove": GClefLigatedNumberAbove,
	"gClefLigatedNumberBelow": GClefLigatedNumberBelow,
	"gClefReversed": GClefReversed,
	"gClefTurned": GClefTurned,
	"glissandoDown": GlissandoDown,
	"glissandoUp": GlissandoUp,
	"graceNoteAcciaccaturaStemDown": GraceNoteAcciaccaturaStemDown,
	"graceNoteAcciaccaturaStemUp": GraceNoteAcciaccaturaStemUp,
	"graceNoteAppoggiaturaStemDown": GraceNoteAppoggiaturaStemDown,
	"graceNoteAppoggiaturaStemUp": GraceNoteAppoggiaturaStemUp,
	"graceNoteSlashStemDown": GraceNoteSlashStemDown,
	"graceNoteSlashStemUp": GraceNoteSlashStemUp,
	"guitarBarreFull": GuitarBarreFull,
	"guitarBarreHalf": GuitarBarreHalf,
	"guitarClosePedal": GuitarClosePedal,
	"guitarFadeIn": GuitarFadeIn,
	"guitarFadeOut": GuitarFadeOut,
	"guitarGolpe": GuitarGolpe,
	"guitarHalfOpenPedal": GuitarHalfOpenPedal,
	"guitarLeftHandTapping": GuitarLeftHandTapping,
	"guitarOpenPedal": GuitarOpenPedal,
	"guitarRightHandTapping": GuitarRightHandTapping,
	"guitarShake": GuitarShake,
	"guitarString0": GuitarString0,
	"guitarString1": GuitarString1,
	"guitarString10": GuitarString10,
	"guitarString11": GuitarString11,
	"guitarString12": GuitarString12,
	"guitarString13": GuitarString13,
	"guitarString2": GuitarString2,
	"guitarString3": GuitarString3,
	"guitarString4": GuitarString4,
	"guitarString5": GuitarString5,
	"guitarString6": GuitarString6,
	"guitarString7": GuitarString7,
	"guitarString8": GuitarString8,
	"guitarString9": GuitarString9,
	"guitarStrumDown": GuitarStrumDown,
	"guitarStrumUp": GuitarStrumUp,
	"guitarVibratoBarDip": GuitarVibratoBarDip,
	"guitarVibratoBarScoop": GuitarVibratoBarScoop,
	"guitarVibratoStroke": GuitarVibratoStroke,
	"guitarVolumeSwell": GuitarVolumeSwell,
	"guitarWideVibratoStroke": GuitarWideVibratoStroke,
	"handbellsBelltree": HandbellsBelltree,
	"handbellsDamp3": HandbellsDamp3,
	"handbellsEcho1": HandbellsEcho1,
	"handbellsEcho2": HandbellsEcho2,
	"handbellsGyro": HandbellsGyro,
	"handbellsHandMartellato": HandbellsHandMartellato,
	"handbellsMalletBellOnTable": HandbellsMalletBellOnTable,
	"handbellsMalletBellSuspended": HandbellsMalletBellSuspended,
	"handbellsMalletLft": HandbellsMalletLft,
	"handbellsMartellato": HandbellsMartellato,
	"handbellsMartellatoLift": HandbellsMartellatoLift,
	"handbellsMutedMartellato": HandbellsMutedMartellato,
	"handbellsPluckLift": HandbellsPluckLift,
	"handbellsSwing": HandbellsSwing,
	"handbellsSwingDown": HandbellsSwingDown,
	"handbellsSwingUp": HandbellsSwingUp,
	"handbellsTablePairBells": HandbellsTablePairBells,
	"handbellsTableSingleBell": HandbellsTableSingleBell,
	"harpMetalRod": HarpMetalRod,
	"harpPedalCentered": HarpPedalCentered,
	"harpPedalDivider": HarpPedalDivider,
	"harpPedalLowered": HarpPedalLowered,
	"harpPedalRaised": HarpPedalRaised,
	"harpSalzedoAeolianAscending": HarpSalzedoAeolianAscending,
	"harpSalzedoAeolianDescending": HarpSalzedoAeolianDescending,
	"harpSalzedoDampAbove": HarpSalzedoDampAbove,
	"harpSalzedoDampBelow": HarpSalzedoDampBelow,
	"harpSalzedoDampBothHands": HarpSalzedoDampBothHands,
	"harpSalzedoDampLowStrings": HarpSalzedoDampLowStrings,
	"harpSalzedoFluidicSoundsLeft": HarpSalzedoFluidicSoundsLeft,
	"harpSalzedoFluidicSoundsRight": HarpSalzedoFluidicSoundsRight,
	"harpSalzedoIsolatedSounds": HarpSalzedoIsolatedSounds,
	"harpSalzedoMetallicSounds": HarpSalzedoMetallicSounds,
	"harpSalzedoMetallicSoundsOneString": HarpSalzedoMetallicSoundsOneString,
	"harpSalzedoMuffleTotally": HarpSalzedoMuffleTotally,
	"harpSalzedoOboicFlux": HarpSalzedoOboicFlux,
	"harpSalzedoPlayUpperEnd": HarpSalzedoPlayUpperEnd,
	"harpSalzedoSlideWithSuppleness": HarpSalzedoSlideWithSuppleness,
	"harpSalzedoSnareDrum": HarpSalzedoSnareDrum,
	"harpSalzedoTamTamSounds": HarpSalzedoTamTamSounds,
	"harpSalzedoThunderEffect": HarpSalzedoThunderEffect,
	"harpSalzedoTimpanicSounds": HarpSalzedoTimpanicSounds,
	"harpSalzedoWhistlingSounds": HarpSalzedoWhistlingSounds,
	"harpStringNoiseStem": HarpStringNoiseStem,
	"harpTuningKey": HarpTuningKey,
	"harpTuningKeyGlissando": HarpTuningKeyGlissando,
	"harpTuningKeyHandle": HarpTuningKeyHandle,
	"harpTuningKeyShank": HarpTuningKeyShank,
	"indianDrumClef": IndianDrumClef,
	"kahnBackChug": KahnBackChug,
	"kahnBackFlap": KahnBackFlap,
	"kahnBackRiff": KahnBackRiff,
	"kahnBackRip": KahnBackRip,
	"kahnBallChange": KahnBallChange,
	"kahnBallDig": KahnBallDig,
	"kahnBrushBackward": KahnBrushBackward,
	"kahnBrushForward": KahnBrushForward,
	"kahnChug": KahnChug,
	"kahnClap": KahnClap,
	"kahnDoubleSnap": KahnDoubleSnap,
	"kahnDoubleWing": KahnDoubleWing,
	"kahnDrawStep": KahnDrawStep,
	"kahnDrawTap": KahnDrawTap,
	"kahnFlam": KahnFlam,
	"kahnFlap": KahnFlap,
	"kahnFlapStep": KahnFlapStep,
	"kahnFlat": KahnFlat,
	"kahnFleaHop": KahnFleaHop,
	"kahnFleaTap": KahnFleaTap,
	"kahnGraceTap": KahnGraceTap,
	"kahnGraceTapChange": KahnGraceTapChange,
	"kahnGraceTapHop": KahnGraceTapHop,
	"kahnGraceTapStamp": KahnGraceTapStamp,
	"kahnHeel": KahnHeel,
	"kahnHeelChange": KahnHeelChange,
	"kahnHeelClick": KahnHeelClick,
	"kahnHeelDrop": KahnHeelDrop,
	"kahnHeelStep": KahnHeelStep,
	"kahnHeelTap": KahnHeelTap,
	"kahnHop": KahnHop,
	"kahnJumpApart": KahnJumpApart,
	"kahnJumpTogether": KahnJumpTogether,
	"kahnKneeInward": KahnKneeInward,
	"kahnKneeOutward": KahnKneeOutward,
	"kahnLeap": KahnLeap,
	"kahnLeapFlatFoot": KahnLeapFlatFoot,
	"kahnLeapHeelClick": KahnLeapHeelClick,
	"kahnLeftCatch": KahnLeftCatch,
	"kahnLeftCross": KahnLeftCross,
	"kahnLeftFoot": KahnLeftFoot,
	"kahnLeftToeStrike": KahnLeftToeStrike,
	"kahnLeftTurn": KahnLeftTurn,
	"kahnOverTheTop": KahnOverTheTop,
	"kahnOverTheTopTap": KahnOverTheTopTap,
	"kahnPull": KahnPull,
	"kahnPush": KahnPush,
	"kahnRiff": KahnRiff,
	"kahnRiffle": KahnRiffle,
	"kahnRightCatch": KahnRightCatch,
	"kahnRightCross": KahnRightCross,
	"kahnRightFoot": KahnRightFoot,
	"kahnRightToeStrike": KahnRightToeStrike,
	"kahnRightTurn": KahnRightTurn,
	"kahnRip": KahnRip,
	"kahnRipple": KahnRipple,
	"kahnScrape": KahnScrape,
	"kahnScuff": KahnScuff,
	"kahnScuffle": KahnScuffle,
	"kahnShuffle": KahnShuffle,
	"kahnSlam": KahnSlam,
	"kahnSlap": KahnSlap,
	"kahnSlideStep": KahnSlideStep,
	"kahnSlideTap": KahnSlideTap,
	"kahnSnap": KahnSnap,
	"kahnStamp": KahnStamp,
	"kahnStampStamp": KahnStampStamp,
	"kahnStep": KahnStep,
	"kahnStepStamp": KahnStepStamp,
	"kahnStomp": KahnStomp,
	"kahnStompBrush": KahnStompBrush,
	"kahnTap": KahnTap,
	"kahnToe": KahnToe,
	"kahnToeClick": KahnToeClick,
	"kahnToeDrop": KahnToeDrop,
	"kahnToeStep": KahnToeStep,
	"kahnToeTap": KahnToeTap,
	"kahnTrench": KahnTrench,
	"kahnWing": KahnWing,
	"kahnWingChange": KahnWingChange,
	"kahnZank": KahnZank,
	"kahnZink": KahnZink,
	"keyboardBebung2DotsAbove": KeyboardBebung2DotsAbove,
	"keyboardBebung2DotsBelow": KeyboardBebung2DotsBelow,
	"keyboardBebung3DotsAbove": KeyboardBebung3DotsAbove,
	"keyboardBebung3DotsBelow": KeyboardBebung3DotsBelow,
	"keyboardBebung4DotsAbove": KeyboardBebung4DotsAbove,
	"keyboardBebung4DotsBelow": KeyboardBebung4DotsBelow,
	"keyboardLeftPedalPictogram": KeyboardLeftPedalPictogram,
	"keyboardMiddlePedalPictogram": KeyboardMiddlePedalPictogram,
	"keyboardPedalD": KeyboardPedalD,
	"keyboardPedalDot": KeyboardPedalDot,
	"keyboardPedalE": KeyboardPedalE,
	"keyboardPedalHalf": KeyboardPedalHalf,
	"keyboardPedalHalf2": KeyboardPedalHalf2,
	"keyboardPedalHalf3": KeyboardPedalHalf3,
	"keyboardPedalHeel1": KeyboardPedalHeel1,
	"keyboardPedalHeel2": KeyboardPedalHeel2,
	"keyboardPedalHeel3": KeyboardPedalHeel3,
	"keyboardPedalHeelToToe": KeyboardPedalHeelToToe,
	"keyboardPedalHeelToe": KeyboardPedalHeelToe,
	"keyboardPedalHookEnd": KeyboardPedalHookEnd,
	"keyboardPedalHookStart": KeyboardPedalHookStart,
	"keyboardPedalHyphen": KeyboardPedalHyphen,
	"keyboardPedalP": KeyboardPedalP,
	"keyboardPedalParensLeft": KeyboardPedalParensLeft,
	"keyboardPedalParensRight": KeyboardPedalParensRight,
	"keyboardPedalPed": KeyboardPedalPed,
	"keyboardPedalS": KeyboardPedalS,
	"keyboardPedalSost": KeyboardPedalSost,
	"keyboardPedalToe1": KeyboardPedalToe1,
	"keyboardPedalToe2": KeyboardPedalToe2,
	"keyboardPedalToeToHeel": KeyboardPedalToeToHeel,
	"keyboardPedalUp": KeyboardPedalUp,
	"keyboardPedalUpNotch": KeyboardPedalUpNotch,
	"keyboardPedalUpSpecial": KeyboardPedalUpSpecial,
	"keyboardPlayWithLH": KeyboardPlayWithLH,
	"keyboardPlayWithLHEnd": KeyboardPlayWithLHEnd,
	"keyboardPlayWithRH": KeyboardPlayWithRH,
	"keyboardPlayWithRHEnd": KeyboardPlayWithRHEnd,
	"keyboardPluckInside": KeyboardPluckInside,
	"keyboardRightPedalPictogram": KeyboardRightPedalPictogram,
	"kievanAccidentalFlat": KievanAccidentalFlat,
	"kievanAccidentalSharp": KievanAccidentalSharp,
	"kievanAugmentationDot": KievanAugmentationDot,
	"kievanCClef": KievanCClef,
	"kievanEndingSymbol": KievanEndingSymbol,
	"kievanNote8thStemDown": KievanNote8thStemDown,
	"kievanNote8thStemUp": KievanNote8thStemUp,
	"kievanNoteBeam": KievanNoteBeam,
	"kievanNoteHalfStaffLine": KievanNoteHalfStaffLine,
	"kievanNoteHalfStaffSpace": KievanNoteHalfStaffSpace,
	"kievanNoteQuarterStemDown": KievanNoteQuarterStemDown,
	"kievanNoteQuarterStemUp": KievanNoteQuarterStemUp,
	"kievanNoteReciting": KievanNoteReciting,
	"kievanNoteWhole": KievanNoteWhole,
	"kievanNoteWholeFinal": KievanNoteWholeFinal,
	"kodalyHandDo": KodalyHandDo,
	"kodalyHandFa": KodalyHandFa,
	"kodalyHandLa": KodalyHandLa,
	"kodalyHandMi": KodalyHandMi,
	"kodalyHandRe": KodalyHandRe,
	"kodalyHandSo": KodalyHandSo,
	"kodalyHandTi": KodalyHandTi,
	"leftRepeatSmall": LeftRepeatSmall,
	"legerLine": LegerLine,
	"legerLineNarrow": LegerLineNarrow,
	"legerLineWide": LegerLineWide,
	"luteBarlineEndRepeat": LuteBarlineEndRepeat,
	"luteBarlineFinal": LuteBarlineFinal,
	"luteBarlineStartRepeat": LuteBarlineStartRepeat,
	"luteDuration16th": LuteDuration16th,
	"luteDuration32nd": LuteDuration32nd,
	"luteDuration8th": LuteDuration8th,
	"luteDurationDoubleWhole": LuteDurationDoubleWhole,
	"luteDurationHalf": LuteDurationHalf,
	"luteDurationQuarter": LuteDurationQuarter,
	"luteDurationWhole": LuteDurationWhole,
	"luteFingeringRHFirst": LuteFingeringRHFirst,
	"luteFingeringRHSecond": LuteFingeringRHSecond,
	"luteFingeringRHThird": LuteFingeringRHThird,
	"luteFingeringRHThumb": LuteFingeringRHThumb,
	"luteFrench10thCourse": LuteFrench10thCourse,
	"luteFrench7thCourse": LuteFrench7thCourse,
	"luteFrench8thCourse": LuteFrench8thCourse,
	"luteFrench9thCourse": LuteFrench9thCourse,
	"luteFrenchAppoggiaturaAbove": LuteFrenchAppoggiaturaAbove,
	"luteFrenchAppoggiaturaBelow": LuteFrenchAppoggiaturaBelow,
	"luteFrenchFretA": LuteFrenchFretA,
	"luteFrenchFretB": LuteFrenchFretB,
	"luteFrenchFretC": LuteFrenchFretC,
	"luteFrenchFretD": LuteFrenchFretD,
	"luteFrenchFretE": LuteFrenchFretE,
	"luteFrenchFretF": LuteFrenchFretF,
	"luteFrenchFretG": LuteFrenchFretG,
	"luteFrenchFretH": LuteFrenchFretH,
	"luteFrenchFretI": LuteFrenchFretI,
	"luteFrenchFretK": LuteFrenchFretK,
	"luteFrenchFretL": LuteFrenchFretL,
	"luteFrenchFretM": LuteFrenchFretM,
	"luteFrenchFretN": LuteFrenchFretN,
	"luteFrenchMordentInverted": LuteFrenchMordentInverted,
	"luteFrenchMordentLower": LuteFrenchMordentLower,
	"luteFrenchMordentUpper": LuteFrenchMordentUpper,
	"luteGermanALower": LuteGermanALower,
	"luteGermanAUpper": LuteGermanAUpper,
	"luteGermanBLower": LuteGermanBLower,
	"luteGermanBUpper": LuteGermanBUpper,
	"luteGermanCLower": LuteGermanCLower,
	"luteGermanCUpper": LuteGermanCUpper,
	"luteGermanDLower": LuteGermanDLower,
	"luteGermanDUpper": LuteGermanDUpper,
	"luteGermanELower": LuteGermanELower,
	"luteGermanEUpper": LuteGermanEUpper,
	"luteGermanFLower": LuteGermanFLower,
	"luteGermanFUpper": LuteGermanFUpper,
	"luteGermanGLower": LuteGermanGLower,
	"luteGermanGUpper": LuteGermanGUpper,
	"luteGermanHLower": LuteGermanHLower,
	"luteGermanHUpper": LuteGermanHUpper,
	"luteGermanILower": LuteGermanILower,
	"luteGermanIUpper": LuteGermanIUpper,
	"luteGermanKLower": LuteGermanKLower,
	"luteGermanKUpper": LuteGermanKUpper,
	"luteGermanLLower": LuteGermanLLower,
	"luteGermanLUpper": LuteGermanLUpper,
	"luteGermanMLower": LuteGermanMLower,
	"luteGermanMUpper": LuteGermanMUpper,
	"luteGermanNLower": LuteGermanNLower,
	"luteGermanNUpper": LuteGermanNUpper,
	"luteGermanOLower": LuteGermanOLower,
	"luteGermanPLower": LuteGermanPLower,
	"luteGermanQLower": LuteGermanQLower,
	"luteGermanRLower": LuteGermanRLower,
	"luteGermanSLower": LuteGermanSLower,
	"luteGermanTLower": LuteGermanTLower,
	"luteGermanVLower": LuteGermanVLower,
	"luteGermanXLower": LuteGermanXLower,
	"luteGermanYLower": LuteGermanYLower,
	"luteGermanZLower": LuteGermanZLower,
	"luteItalianClefCSolFaUt": LuteItalianClefCSolFaUt,
	"luteItalianClefFFaUt": LuteItalianClefFFaUt,
	"luteItalianFret0": LuteItalianFret0,
	"luteItalianFret1": LuteItalianFret1,
	"luteItalianFret2": LuteItalianFret2,
	"luteItalianFret3": LuteItalianFret3,
	"luteItalianFret4": LuteItalianFret4,
	"luteItalianFret5": LuteItalianFret5,
	"luteItalianFret6": LuteItalianFret6,
	"luteItalianFret7": LuteItalianFret7,
	"luteItalianFret8": LuteItalianFret8,
	"luteItalianFret9": LuteItalianFret9,
	"luteItalianHoldFinger": LuteItalianHoldFinger,
	"luteItalianHoldNote": LuteItalianHoldNote,
	"luteItalianReleaseFinger": LuteItalianReleaseFinger,
	"luteItalianTempoFast": LuteItalianTempoFast,
	"luteItalianTempoNeitherFastNorSlow": LuteItalianTempoNeitherFastNorSlow,
	"luteItalianTempoSlow": LuteItalianTempoSlow,
	"luteItalianTempoSomewhatFast": LuteItalianTempoSomewhatFast,
	"luteItalianTempoVerySlow": LuteItalianTempoVerySlow,
	"luteItalianTimeTriple": LuteItalianTimeTriple,
	"luteItalianTremolo": LuteItalianTremolo,
	"luteItalianVibrato": LuteItalianVibrato,
	"luteStaff6Lines": LuteStaff6Lines,
	"luteStaff6LinesNarrow": LuteStaff6LinesNarrow,
	"luteStaff6LinesWide": LuteStaff6LinesWide,
	"lyricsElision": LyricsElision,
	"lyricsElisionNarrow": LyricsElisionNarrow,
	"lyricsElisionWide": LyricsElisionWide,
	"lyricsHyphenBaseline": LyricsHyphenBaseline,
	"lyricsHyphenBaselineNonBreaking": LyricsHyphenBaselineNonBreaking,
	"lyricsTextRepeat": LyricsTextRepeat,
	"medRenFlatHardB": MedRenFlatHardB,
	"medRenFlatSoftB": MedRenFlatSoftB,
	"medRenFlatWithDot": MedRenFlatWithDot,
	"medRenGClefCMN": MedRenGClefCMN,
	"medRenLiquescenceCMN": MedRenLiquescenceCMN,
	"medRenLiquescentAscCMN": MedRenLiquescentAscCMN,
	"medRenLiquescentDescCMN": MedRenLiquescentDescCMN,
	"medRenNatural": MedRenNatural,
	"medRenNaturalWithCross": MedRenNaturalWithCross,
	"medRenOriscusCMN": MedRenOriscusCMN,
	"medRenPlicaCMN": MedRenPlicaCMN,
	"medRenPunctumCMN": MedRenPunctumCMN,
	"medRenQuilismaCMN": MedRenQuilismaCMN,
	"medRenSharpCroix": MedRenSharpCroix,
	"medRenStrophicusCMN": MedRenStrophicusCMN,
	"mensuralAlterationSign": MensuralAlterationSign,
	"mensuralBlackBrevis": MensuralBlackBrevis,
	"mensuralBlackBrevisVoid": MensuralBlackBrevisVoid,
	"mensuralBlackDragma": MensuralBlackDragma,
	"mensuralBlackLonga": MensuralBlackLonga,
	"mensuralBlackMaxima": MensuralBlackMaxima,
	"mensuralBlackMinima": MensuralBlackMinima,
	"mensuralBlackMinimaVoid": MensuralBlackMinimaVoid,
	"mensuralBlackSemibrevis": MensuralBlackSemibrevis,
	"mensuralBlackSemibrevisCaudata": MensuralBlackSemibrevisCaudata,
	"mensuralBlackSemibrevisOblique": MensuralBlackSemibrevisOblique,
	"mensuralBlackSemibrevisVoid": MensuralBlackSemibrevisVoid,
	"mensuralBlackSemiminima": MensuralBlackSemiminima,
	"mensuralCclef": MensuralCclef,
	"mensuralCclefPetrucciPosHigh": MensuralCclefPetrucciPosHigh,
	"mensuralCclefPetrucciPosHighest": MensuralCclefPetrucciPosHighest,
	"mensuralCclefPetrucciPosLow": MensuralCclefPetrucciPosLow,
	"mensuralCclefPetrucciPosLowest": MensuralCclefPetrucciPosLowest,
	"mensuralCclefPetrucciPosMiddle": MensuralCclefPetrucciPosMiddle,
	"mensuralColorationEndRound": MensuralColorationEndRound,
	"mensuralColorationEndSquare": MensuralColorationEndSquare,
	"mensuralColorationStartRound": MensuralColorationStartRound,
	"mensuralColorationStartSquare": MensuralColorationStartSquare,
	"mensuralCombStemDiagonal": MensuralCombStemDiagonal,
	"mensuralCombStemDown": MensuralCombStemDown,
	"mensuralCombStemDownFlagExtended": MensuralCombStemDownFlagExtended,
	"mensuralCombStemDownFlagFlared": MensuralCombStemDownFlagFlared,
	"mensuralCombStemDownFlagFusa": MensuralCombStemDownFlagFusa,
	"mensuralCombStemDownFlagLeft": MensuralCombStemDownFlagLeft,
	"mensuralCombStemDownFlagRight": MensuralCombStemDownFlagRight,
	"mensuralCombStemDownFlagSemiminima": MensuralCombStemDownFlagSemiminima,
	"mensuralCombStemUp": MensuralCombStemUp,
	"mensuralCombStemUpFlagExtended": MensuralCombStemUpFlagExtended,
	"mensuralCombStemUpFlagFlared": MensuralCombStemUpFlagFlared,
	"mensuralCombStemUpFlagFusa": MensuralCombStemUpFlagFusa,
	"mensuralCombStemUpFlagLeft": MensuralCombStemUpFlagLeft,
	"mensuralCombStemUpFlagRight": MensuralCombStemUpFlagRight,
	"mensuralCombStemUpFlagSemiminima": MensuralCombStemUpFlagSemiminima,
	"mensuralCustosCheckmark": MensuralCustosCheckmark,
	"mensuralCustosDown": MensuralCustosDown,
	"mensuralCustosTurn": MensuralCustosTurn,
	"mensuralCustosUp": MensuralCustosUp,
	"mensuralFclef": MensuralFclef,
	"mensuralFclefPetrucci": MensuralFclefPetrucci,
	"mensuralGclef": MensuralGclef,
	"mensuralGclefPetrucci": MensuralGclefPetrucci,
	"mensuralModusImperfectumVert": MensuralModusImperfectumVert,
	"mensuralModusPerfectumVert": MensuralModusPerfectumVert,
	"mensuralNoteheadLongaBlack": MensuralNoteheadLongaBlack,
	"mensuralNoteheadLongaBlackVoid": MensuralNoteheadLongaBlackVoid,
	"mensuralNoteheadLongaVoid": MensuralNoteheadLongaVoid,
	"mensuralNoteheadLongaWhite": MensuralNoteheadLongaWhite,
	"mensuralNoteheadMaximaBlack": MensuralNoteheadMaximaBlack,
	"mensuralNoteheadMaximaBlackVoid": MensuralNoteheadMaximaBlackVoid,
	"mensuralNoteheadMaximaVoid": MensuralNoteheadMaximaVoid,
	"mensuralNoteheadMaximaWhite": MensuralNoteheadMaximaWhite,
	"mensuralNoteheadMinimaWhite": MensuralNoteheadMinimaWhite,
	"mensuralNoteheadSemibrevisBlack": MensuralNoteheadSemibrevisBlack,
	"mensuralNoteheadSemibrevisBlackVoid": MensuralNoteheadSemibrevisBlackVoid,
	"mensuralNoteheadSemibrevisBlackVoidTurned": MensuralNoteheadSemibrevisBlackVoidTurned,
	"mensuralNoteheadSemibrevisVoid": MensuralNoteheadSemibrevisVoid,
	"mensuralNoteheadSemiminimaWhite": MensuralNoteheadSemiminimaWhite,
	"mensuralObliqueAsc2ndBlack": MensuralObliqueAsc2ndBlack,
	"mensuralObliqueAsc2ndBlackVoid": MensuralObliqueAsc2ndBlackVoid,
	"mensuralObliqueAsc2ndVoid": MensuralObliqueAsc2ndVoid,
	"mensuralObliqueAsc2ndWhite": MensuralObliqueAsc2ndWhite,
	"mensuralObliqueAsc3rdBlack": MensuralObliqueAsc3rdBlack,
	"mensuralObliqueAsc3rdBlackVoid": MensuralObliqueAsc3rdBlackVoid,
	"mensuralObliqueAsc3rdVoid": MensuralObliqueAsc3rdVoid,
	"mensuralObliqueAsc3rdWhite": MensuralObliqueAsc3rdWhite,
	"mensuralObliqueAsc4thBlack": MensuralObliqueAsc4thBlack,
	"mensuralObliqueAsc4thBlackVoid": MensuralObliqueAsc4thBlackVoid,
	"mensuralObliqueAsc4thVoid": MensuralObliqueAsc4thVoid,
	"mensuralObliqueAsc4thWhite": MensuralObliqueAsc4thWhite,
	"mensuralObliqueAsc5thBlack": MensuralObliqueAsc5thBlack,
	"mensuralObliqueAsc5thBlackVoid": MensuralObliqueAsc5thBlackVoid,
	"mensuralObliqueAsc5thVoid": MensuralObliqueAsc5thVoid,
	"mensuralObliqueAsc5thWhite": MensuralObliqueAsc5thWhite,
	"mensuralObliqueDesc2ndBlack": MensuralObliqueDesc2ndBlack,
	"mensuralObliqueDesc2ndBlackVoid": MensuralObliqueDesc2ndBlackVoid,
	"mensuralObliqueDesc2ndVoid": MensuralObliqueDesc2ndVoid,
	"mensuralObliqueDesc2ndWhite": MensuralObliqueDesc2ndWhite,
	"mensuralObliqueDesc3rdBlack": MensuralObliqueDesc3rdBlack,
	"mensuralObliqueDesc3rdBlackVoid": MensuralObliqueDesc3rdBlackVoid,
	"mensuralObliqueDesc3rdVoid": MensuralObliqueDesc3rdVoid,
	"mensuralObliqueDesc3rdWhite": MensuralObliqueDesc3rdWhite,
	"mensuralObliqueDesc4thBlack": MensuralObliqueDesc4thBlack,
	"mensuralObliqueDesc4thBlackVoid": MensuralObliqueDesc4thBlackVoid,
	"mensuralObliqueDesc4thVoid": MensuralObliqueDesc4thVoid,
	"mensuralObliqueDesc4thWhite": MensuralObliqueDesc4thWhite,
	"mensuralObliqueDesc5thBlack": MensuralObliqueDesc5thBlack,
	"mensuralObliqueDesc5thBlackVoid": MensuralObliqueDesc5thBlackVoid,
	"mensuralObliqueDesc5thVoid": MensuralObliqueDesc5thVoid,
	"mensuralObliqueDesc5thWhite": MensuralObliqueDesc5thWhite,
	"mensuralProlation1": MensuralProlation1,
	"mensuralProlation10": MensuralProlation10,
	"mensuralProlation11": MensuralProlation11,
	"mensuralProlation2": MensuralProlation2,
	"mensuralProlation3": MensuralProlation3,
	"mensuralProlation4": MensuralProlation4,
	"mensuralProlation5": MensuralProlation5,
	"mensuralProlation6": MensuralProlation6,
	"mensuralProlation7": MensuralProlation7,
	"mensuralProlation8": MensuralProlation8,
	"mensuralProlation9": MensuralProlation9,
	"mensuralProlationCombiningDot": MensuralProlationCombiningDot,
	"mensuralProlationCombiningDotVoid": MensuralProlationCombiningDotVoid,
	"mensuralProlationCombiningStroke": MensuralProlationCombiningStroke,
	"mensuralProlationCombiningThreeDots": MensuralProlationCombiningThreeDots,
	"mensuralProlationCombiningThreeDotsTri": MensuralProlationCombiningThreeDotsTri,
	"mensuralProlationCombiningTwoDots": MensuralProlationCombiningTwoDots,
	"mensuralProportion1": MensuralProportion1,
	"mensuralProportion2": MensuralProportion2,
	"mensuralProportion3": MensuralProportion3,
	"mensuralProportion4": MensuralProportion4,
	"mensuralProportion5": MensuralProportion5,
	"mensuralProportion6": MensuralProportion6,
	"mensuralProportion7": MensuralProportion7,
	"mensuralProportion8": MensuralProportion8,
	"mensuralProportion9": MensuralProportion9,
	"mensuralProportionMajor": MensuralProportionMajor,
	"mensuralProportionMinor": MensuralProportionMinor,
	"mensuralProportionProportioDupla1": MensuralProportionProportioDupla1,
	"mensuralProportionProportioDupla2": MensuralProportionProportioDupla2,
	"mensuralProportionProportioQuadrupla": MensuralProportionProportioQuadrupla,
	"mensuralProportionProportioTripla": MensuralProportionProportioTripla,
	"mensuralProportionTempusPerfectum": MensuralProportionTempusPerfectum,
	"mensuralRestBrevis": MensuralRestBrevis,
	"mensuralRestFusa": MensuralRestFusa,
	"mensuralRestLongaImperfecta": MensuralRestLongaImperfecta,
	"mensuralRestLongaPerfecta": MensuralRestLongaPerfecta,
	"mensuralRestMaxima": MensuralRestMaxima,
	"mensuralRestMinima": MensuralRestMinima,
	"mensuralRestSemibrevis": MensuralRestSemibrevis,
	"mensuralRestSemifusa": MensuralRestSemifusa,
	"mensuralRestSemiminima": MensuralRestSemiminima,
	"mensuralSignumDown": MensuralSignumDown,
	"mensuralSignumUp": MensuralSignumUp,
	"mensuralTempusImperfectumHoriz": MensuralTempusImperfectumHoriz,
	"mensuralTempusPerfectumHoriz": MensuralTempusPerfectumHoriz,
	"mensuralWhiteBrevis": MensuralWhiteBrevis,
	"mensuralWhiteFusa": MensuralWhiteFusa,
	"mensuralWhiteLonga": MensuralWhiteLonga,
	"mensuralWhiteMaxima": MensuralWhiteMaxima,
	"mensuralWhiteMinima": MensuralWhiteMinima,
	"mensuralWhiteSemibrevis": MensuralWhiteSemibrevis,
	"mensuralWhiteSemiminima": MensuralWhiteSemiminima,
	"metAugmentationDot": MetAugmentationDot,
	"metNote1024thDown": MetNote1024thDown,
	"metNote1024thUp": MetNote1024thUp,
	"metNote128thDown": MetNote128thDown,
	"metNote128thUp": MetNote128thUp,
	"metNote16thDown": MetNote16thDown,
	"metNote16thUp": MetNote16thUp,
	"metNote256thDown": MetNote256thDown,
	"metNote256thUp": MetNote256thUp,
	"metNote32ndDown": MetNote32ndDown,
	"metNote32ndUp": MetNote32ndUp,
	"metNote512thDown": MetNote512thDown,
	"metNote512thUp": MetNote512thUp,
	"metNote64thDown": MetNote64thDown,
	"metNote64thUp": MetNote64thUp,
	"metNote8thDown": MetNote8thDown,
	"metNote8thUp": MetNote8thUp,
	"metNoteDoubleWhole": MetNoteDoubleWhole,
	"metNoteDoubleWholeSquare": MetNoteDoubleWholeSquare,
	"metNoteHalfDown": MetNoteHalfDown,
	"metNoteHalfUp": MetNoteHalfUp,
	"metNoteQuarterDown": MetNoteQuarterDown,
	"metNoteQuarterUp": MetNoteQuarterUp,
	"metNoteWhole": MetNoteWhole,
	"metricModulationArrowLeft": MetricModulationArrowLeft,
	"metricModulationArrowRight": MetricModulationArrowRight,
	"miscDoNotCopy": MiscDoNotCopy,
	"miscDoNotPhotocopy": MiscDoNotPhotocopy,
	"miscEyeglasses": MiscEyeglasses,
	"note1024thDown": Note1024thDown,
	"note1024thUp": Note1024thUp,
	"note128thDown": Note128thDown,
	"note128thUp": Note128thUp,
	"note16thDown": Note16thDown,
	"note16thUp": Note16thUp,
	"note256thDown": Note256thDown,
	"note256thUp": Note256thUp,
	"note32ndDown": Note32ndDown,
	"note32ndUp": Note32ndUp,
	"note512thDown": Note512thDown,
	"note512thUp": Note512thUp,
	"note64thDown": Note64thDown,
	"note64thUp": Note64thUp,
	"note8thDown": Note8thDown,
	"note8thUp": Note8thUp,
	"noteABlack": NoteABlack,
	"noteAFlatBlack": NoteAFlatBlack,
	"noteAFlatHalf": NoteAFlatHalf,
	"noteAFlatWhole": NoteAFlatWhole,
	"noteAHalf": NoteAHalf,
	"noteASharpBlack": NoteASharpBlack,
	"noteASharpHalf": NoteASharpHalf,
	"noteASharpWhole": NoteASharpWhole,
	"noteAWhole": NoteAWhole,
	"noteBBlack": NoteBBlack,
	"noteBFlatBlack": NoteBFlatBlack,
	"noteBFlatHalf": NoteBFlatHalf,
	"noteBFlatWhole": NoteBFlatWhole,
	"noteBHalf": NoteBHalf,
	"noteBSharpBlack": NoteBSharpBlack,
	"noteBSharpHalf": NoteBSharpHalf,
	"noteBSharpWhole": NoteBSharpWhole,
	"noteBWhole": NoteBWhole,
	"noteCBlack": NoteCBlack,
	"noteCFlatBlack": NoteCFlatBlack,
	"noteCFlatHalf": NoteCFlatHalf,
	"noteCFlatWhole": NoteCFlatWhole,
	"noteCHalf": NoteCHalf,
	"noteCSharpBlack": NoteCSharpBlack,
	"noteCSharpHalf": NoteCSharpHalf,
	"noteCSharpWhole": NoteCSharpWhole,
	"noteCWhole": NoteCWhole,
	"noteDBlack": NoteDBlack,
	"noteDFlatBlack": NoteDFlatBlack,
	"noteDFlatHalf": NoteDFlatHalf,
	"noteDFlatWhole": NoteDFlatWhole,
	"noteDHalf": NoteDHalf,
	"noteDSharpBlack": NoteDSharpBlack,
	"noteDSharpHalf": NoteDSharpHalf,
	"noteDSharpWhole": NoteDSharpWhole,
	"noteDWhole": NoteDWhole,
	"noteDiBlack": NoteDiBlack,
	"noteDiHalf": NoteDiHalf,
	"noteDiWhole": NoteDiWhole,
	"noteDoBlack": NoteDoBlack,
	"noteDoHalf": NoteDoHalf,
	"noteDoWhole": NoteDoWhole,
	"noteDoubleWhole": NoteDoubleWhole,
	"noteDoubleWholeSquare": NoteDoubleWholeSquare,
	"noteEBlack": NoteEBlack,
	"noteEFlatBlack": NoteEFlatBlack,
	"noteEFlatHalf": NoteEFlatHalf,
	"noteEFlatWhole": NoteEFlatWhole,
	"noteEHalf": NoteEHalf,
	"noteESharpBlack": NoteESharpBlack,
	"noteESharpHalf": NoteESharpHalf,
	"noteESharpWhole": NoteESharpWhole,
	"noteEWhole": NoteEWhole,
	"noteEmptyBlack": NoteEmptyBlack,
	"noteEmptyHalf": NoteEmptyHalf,
	"noteEmptyWhole": NoteEmptyWhole,
	"noteFBlack": NoteFBlack,
	"noteFFlatBlack": NoteFFlatBlack,
	"noteFFlatHalf": NoteFFlatHalf,
	"noteFFlatWhole": NoteFFlatWhole,
	"noteFHalf": NoteFHalf,
	"noteFSharpBlack": NoteFSharpBlack,
	"noteFSharpHalf": NoteFSharpHalf,
	"noteFSharpWhole": NoteFSharpWhole,
	"noteFWhole": NoteFWhole,
	"noteFaBlack": NoteFaBlack,
	"noteFaHalf": NoteFaHalf,
	"noteFaWhole": NoteFaWhole,
	"noteFiBlack": NoteFiBlack,
	"noteFiHalf": NoteFiHalf,
	"noteFiWhole": NoteFiWhole,
	"noteGBlack": NoteGBlack,
	"noteGFlatBlack": NoteGFlatBlack,
	"noteGFlatHalf": NoteGFlatHalf,
	"noteGFlatWhole": NoteGFlatWhole,
	"noteGHalf": NoteGHalf,
	"noteGSharpBlack": NoteGSharpBlack,
	"noteGSharpHalf": NoteGSharpHalf,
	"noteGSharpWhole": NoteGSharpWhole,
	"noteGWhole": NoteGWhole,
	"noteHBlack": NoteHBlack,
	"noteHHalf": NoteHHalf,
	"noteHSharpBlack": NoteHSharpBlack,
	"noteHSharpHalf": NoteHSharpHalf,
	"noteHSharpWhole": NoteHSharpWhole,
	"noteHWhole": NoteHWhole,
	"noteHalfDown": NoteHalfDown,
	"noteHalfUp": NoteHalfUp,
	"noteLaBlack": NoteLaBlack,
	"noteLaHalf": NoteLaHalf,
	"noteLaWhole": NoteLaWhole,
	"noteLeBlack": NoteLeBlack,
	"noteLeHalf": NoteLeHalf,
	"noteLeWhole": NoteLeWhole,
	"noteLiBlack": NoteLiBlack,
	"noteLiHalf": NoteLiHalf,
	"noteLiWhole": NoteLiWhole,
	"noteMeBlack": NoteMeBlack,
	"noteMeHalf": NoteMeHalf,
	"noteMeWhole": NoteMeWhole,
	"noteMiBlack": NoteMiBlack,
	"noteMiHalf": NoteMiHalf,
	"noteMiWhole": NoteMiWhole,
	"noteQuarterDown": NoteQuarterDown,
	"noteQuarterUp": NoteQuarterUp,
	"noteRaBlack": NoteRaBlack,
	"noteRaHalf": NoteRaHalf,
	"noteRaWhole": NoteRaWhole,
	"noteReBlack": NoteReBlack,
	"noteReHalf": NoteReHalf,
	"noteReWhole": NoteReWhole,
	"noteRiBlack": NoteRiBlack,
	"noteRiHalf": NoteRiHalf,
	"noteRiWhole": NoteRiWhole,
	"noteSeBlack": NoteSeBlack,
	"noteSeHalf": NoteSeHalf,
	"noteSeWhole": NoteSeWhole,
	"noteShapeArrowheadLeftBlack": NoteShapeArrowheadLeftBlack,
	"noteShapeArrowheadLeftDoubleWhole": NoteShapeArrowheadLeftDoubleWhole,
	"noteShapeArrowheadLeftWhite": NoteShapeArrowheadLeftWhite,
	"noteShapeDiamondBlack": NoteShapeDiamondBlack,
	"noteShapeDiamondDoubleWhole": NoteShapeDiamondDoubleWhole,
	"noteShapeDiamondWhite": NoteShapeDiamondWhite,
	"noteShapeIsoscelesTriangleBlack": NoteShapeIsoscelesTriangleBlack,
	"noteShapeIsoscelesTriangleDoubleWhole": NoteShapeIsoscelesTriangleDoubleWhole,
	"noteShapeIsoscelesTriangleWhite": NoteShapeIsoscelesTriangleWhite,
	"noteShapeKeystoneBlack": NoteShapeKeystoneBlack,
	"noteShapeKeystoneDoubleWhole": NoteShapeKeystoneDoubleWhole,
	"noteShapeKeystoneWhite": NoteShapeKeystoneWhite,
	"noteShapeMoonBlack": NoteShapeMoonBlack,
	"noteShapeMoonDoubleWhole": NoteShapeMoonDoubleWhole,
	"noteShapeMoonLeftBlack": NoteShapeMoonLeftBlack,
	"noteShapeMoonLeftDoubleWhole": NoteShapeMoonLeftDoubleWhole,
	"noteShapeMoonLeftWhite": NoteShapeMoonLeftWhite,
	"noteShapeMoonWhite": NoteShapeMoonWhite,
	"noteShapeQuarterMoonBlack": NoteShapeQuarterMoonBlack,
	"noteShapeQuarterMoonDoubleWhole": NoteShapeQuarterMoonDoubleWhole,
	"noteShapeQuarterMoonWhite": NoteShapeQuarterMoonWhite,
	"noteShapeRoundBlack": NoteShapeRoundBlack,
	"noteShapeRoundDoubleWhole": NoteShapeRoundDoubleWhole,
	"noteShapeRoundWhite": NoteShapeRoundWhite,
	"noteShapeSquareBlack": NoteShapeSquareBlack,
	"noteShapeSquareDoubleWhole": NoteShapeSquareDoubleWhole,
	"noteShapeSquareWhite": NoteShapeSquareWhite,
	"noteShapeTriangleLeftBlack": NoteShapeTriangleLeftBlack,
	"noteShapeTriangleLeftDoubleWhole": NoteShapeTriangleLeftDoubleWhole,
	"noteShapeTriangleLeftWhite": NoteShapeTriangleLeftWhite,
	"noteShapeTriangleRightBlack": NoteShapeTriangleRightBlack,
	"noteShapeTriangleRightDoubleWhole": NoteShapeTriangleRightDoubleWhole,
	"noteShapeTriangleRightWhite": NoteShapeTriangleRightWhite,
	"noteShapeTriangleRoundBlack": NoteShapeTriangleRoundBlack,
	"noteShapeTriangleRoundDoubleWhole": NoteShapeTriangleRoundDoubleWhole,
	"noteShapeTriangleRoundLeftBlack": NoteShapeTriangleRoundLeftBlack,
	"noteShapeTriangleRoundLeftDoubleWhole": NoteShapeTriangleRoundLeftDoubleWhole,
	"noteShapeTriangleRoundLeftWhite": NoteShapeTriangleRoundLeftWhite,
	"noteShapeTriangleRoundWhite": NoteShapeTriangleRoundWhite,
	"noteShapeTriangleUpBlack": NoteShapeTriangleUpBlack,
	"noteShapeTriangleUpDoubleWhole": NoteShapeTriangleUpDoubleWhole,
	"noteShapeTriangleUpWhite": NoteShapeTriangleUpWhite,
	"noteSiBlack": NoteSiBlack,
	"noteSiHalf": NoteSiHalf,
	"noteSiWhole": NoteSiWhole,
	"noteSoBlack": NoteSoBlack,
	"noteSoHalf": NoteSoHalf,
	"noteSoWhole": NoteSoWhole,
	"noteTeBlack": NoteTeBlack,
	"noteTeHalf": NoteTeHalf,
	"noteTeWhole": NoteTeWhole,
	"noteTiBlack": NoteTiBlack,
	"noteTiHalf": NoteTiHalf,
	"noteTiWhole": NoteTiWhole,
	"noteWhole": NoteWhole,
	"noteheadBlack": NoteheadBlack,
	"noteheadCircleSlash": NoteheadCircleSlash,
	"noteheadCircleX": NoteheadCircleX,
	"noteheadCircleXDoubleWhole": NoteheadCircleXDoubleWhole,
	"noteheadCircleXHalf": NoteheadCircleXHalf,
	"noteheadCircleXWhole": NoteheadCircleXWhole,
	"noteheadCircledBlack": NoteheadCircledBlack,
	"noteheadCircledBlackLarge": NoteheadCircledBlackLarge,
	"noteheadCircledDoubleWhole": NoteheadCircledDoubleWhole,
	"noteheadCircledDoubleWholeLarge": NoteheadCircledDoubleWholeLarge,
	"noteheadCircledHalf": NoteheadCircledHalf,
	"noteheadCircledHalfLarge": NoteheadCircledHalfLarge,
	"noteheadCircledWhole": NoteheadCircledWhole,
	"noteheadCircledWholeLarge": NoteheadCircledWholeLarge,
	"noteheadCircledXLarge": NoteheadCircledXLarge,
	"noteheadClusterDoubleWhole2nd": NoteheadClusterDoubleWhole2nd,
	"noteheadClusterDoubleWhole3rd": NoteheadClusterDoubleWhole3rd,
	"noteheadClusterDoubleWholeBottom": NoteheadClusterDoubleWholeBottom,
	"noteheadClusterDoubleWholeMiddle": NoteheadClusterDoubleWholeMiddle,
	"noteheadClusterDoubleWholeTop": NoteheadClusterDoubleWholeTop,
	"noteheadClusterHalf2nd": NoteheadClusterHalf2nd,
	"noteheadClusterHalf3rd": NoteheadClusterHalf3rd,
	"noteheadClusterHalfBottom": NoteheadClusterHalfBottom,
	"noteheadClusterHalfMiddle": NoteheadClusterHalfMiddle,
	"noteheadClusterHalfTop": NoteheadClusterHalfTop,
	"noteheadClusterQuarter2nd": NoteheadClusterQuarter2nd,
	"noteheadClusterQuarter3rd": NoteheadClusterQuarter3rd,
	"noteheadClusterQuarterBottom": NoteheadClusterQuarterBottom,
	"noteheadClusterQuarterMiddle": NoteheadClusterQuarterMiddle,
	"noteheadClusterQuarterTop": NoteheadClusterQuarterTop,
	"noteheadClusterRoundBlack": NoteheadClusterRoundBlack,
	"noteheadClusterRoundWhite": NoteheadClusterRoundWhite,
	"noteheadClusterSquareBlack": NoteheadClusterSquareBlack,
	"noteheadClusterSquareWhite": NoteheadClusterSquareWhite,
	"noteheadClusterWhole2nd": NoteheadClusterWhole2nd,
	"noteheadClusterWhole3rd": NoteheadClusterWhole3rd,
	"noteheadClusterWholeBottom": NoteheadClusterWholeBottom,
	"noteheadClusterWholeMiddle": NoteheadClusterWholeMiddle,
	"noteheadClusterWholeTop": NoteheadClusterWholeTop,
	"noteheadCowellEleventhNoteSeriesHalf": NoteheadCowellEleventhNoteSeriesHalf,
	"noteheadCowellEleventhNoteSeriesWhole": NoteheadCowellEleventhNoteSeriesWhole,
	"noteheadCowellEleventhSeriesBlack": NoteheadCowellEleventhSeriesBlack,
	"noteheadCowellFifteenthNoteSeriesBlack": NoteheadCowellFifteenthNoteSeriesBlack,
	"noteheadCowellFifteenthNoteSeriesHalf": NoteheadCowellFifteenthNoteSeriesHalf,
	"noteheadCowellFifteenthNoteSeriesWhole": NoteheadCowellFifteenthNoteSeriesWhole,
	"noteheadCowellFifthNoteSeriesBlack": NoteheadCowellFifthNoteSeriesBlack,
	"noteheadCowellFifthNoteSeriesHalf": NoteheadCowellFifthNoteSeriesHalf,
	"noteheadCowellFifthNoteSeriesWhole": NoteheadCowellFifthNoteSeriesWhole,
	"noteheadCowellNinthNoteSeriesBlack": NoteheadCowellNinthNoteSeriesBlack,
	"noteheadCowellNinthNoteSeriesHalf": NoteheadCowellNinthNoteSeriesHalf,
	"noteheadCowellNinthNoteSeriesWhole": NoteheadCowellNinthNoteSeriesWhole,
	"noteheadCowellSeventhNoteSeriesBlack": NoteheadCowellSeventhNoteSeriesBlack,
	"noteheadCowellSeventhNoteSeriesHalf": NoteheadCowellSeventhNoteSeriesHalf,
	"noteheadCowellSeventhNoteSeriesWhole": NoteheadCowellSeventhNoteSeriesWhole,
	"noteheadCowellThirdNoteSeriesBlack": NoteheadCowellThirdNoteSeriesBlack,
	"noteheadCowellThirdNoteSeriesHalf": NoteheadCowellThirdNoteSeriesHalf,
	"noteheadCowellThirdNoteSeriesWhole": NoteheadCowellThirdNoteSeriesWhole,
	"noteheadCowellThirteenthNoteSeriesBlack": NoteheadCowellThirteenthNoteSeriesBlack,
	"noteheadCowellThirteenthNoteSeriesHalf": NoteheadCowellThirteenthNoteSeriesHalf,
	"noteheadCowellThirteenthNoteSeriesWhole": NoteheadCowellThirteenthNoteSeriesWhole,
	"noteheadDiamondBlack": NoteheadDiamondBlack,
	"noteheadDiamondBlackOld": NoteheadDiamondBlackOld,
	"noteheadDiamondBlackWide": NoteheadDiamondBlackWide,
	"noteheadDiamondClusterBlack2nd": NoteheadDiamondClusterBlack2nd,
	"noteheadDiamondClusterBlack3rd": NoteheadDiamondClusterBlack3rd,
	"noteheadDiamondClusterBlackBottom": NoteheadDiamondClusterBlackBottom,
	"noteheadDiamondClusterBlackMiddle": NoteheadDiamondClusterBlackMiddle,
	"noteheadDiamondClusterBlackTop": NoteheadDiamondClusterBlackTop,
	"noteheadDiamondClusterWhite2nd": NoteheadDiamondClusterWhite2nd,
	"noteheadDiamondClusterWhite3rd": NoteheadDiamondClusterWhite3rd,
	"noteheadDiamondClusterWhiteBottom": NoteheadDiamondClusterWhiteBottom,
	"noteheadDiamondClusterWhiteMiddle": NoteheadDiamondClusterWhiteMiddle,
	"noteheadDiamondClusterWhiteTop": NoteheadDiamondClusterWhiteTop,
	"noteheadDiamondDoubleWhole": NoteheadDiamondDoubleWhole,
	"noteheadDiamondDoubleWholeOld": NoteheadDiamondDoubleWholeOld,
	"noteheadDiamondHalf": NoteheadDiamondHalf,
	"noteheadDiamondHalfFilled": NoteheadDiamondHalfFilled,
	"noteheadDiamondHalfOld": NoteheadDiamondHalfOld,
	"noteheadDiamondHalfWide": NoteheadDiamondHalfWide,
	"noteheadDiamondOpen": NoteheadDiamondOpen,
	"noteheadDiamondWhite": NoteheadDiamondWhite,
	"noteheadDiamondWhiteWide": NoteheadDiamondWhiteWide,
	"noteheadDiamondWhole": NoteheadDiamondWhole,
	"noteheadDiamondWholeOld": NoteheadDiamondWholeOld,
	"noteheadDoubleWhole": NoteheadDoubleWhole,
	"noteheadDoubleWholeSquare": NoteheadDoubleWholeSquare,
	"noteheadDoubleWholeWithX": NoteheadDoubleWholeWithX,
	"noteheadHalf": NoteheadHalf,
	"noteheadHalfFilled": NoteheadHalfFilled,
	"noteheadHalfWithX": NoteheadHalfWithX,
	"noteheadHeavyX": NoteheadHeavyX,
	"noteheadHeavyXHat": NoteheadHeavyXHat,
	"noteheadLargeArrowDownBlack": NoteheadLargeArrowDownBlack,
	"noteheadLargeArrowDownDoubleWhole": NoteheadLargeArrowDownDoubleWhole,
	"noteheadLargeArrowDownHalf": NoteheadLargeArrowDownHalf,
	"noteheadLargeArrowDownWhole": NoteheadLargeArrowDownWhole,
	"noteheadLargeArrowUpBlack": NoteheadLargeArrowUpBlack,
	"noteheadLargeArrowUpDoubleWhole": NoteheadLargeArrowUpDoubleWhole,
	"noteheadLargeArrowUpHalf": NoteheadLargeArrowUpHalf,
	"noteheadLargeArrowUpWhole": NoteheadLargeArrowUpWhole,
	"noteheadMoonBlack": NoteheadMoonBlack,
	"noteheadMoonWhite": NoteheadMoonWhite,
	"noteheadNancarrowSine": NoteheadNancarrowSine,
	"noteheadNull": NoteheadNull,
	"noteheadParenthesis": NoteheadParenthesis,
	"noteheadParenthesisLeft": NoteheadParenthesisLeft,
	"noteheadParenthesisRight": NoteheadParenthesisRight,
	"noteheadPlusBlack": NoteheadPlusBlack,
	"noteheadPlusDoubleWhole": NoteheadPlusDoubleWhole,
	"noteheadPlusHalf": NoteheadPlusHalf,
	"noteheadPlusWhole": NoteheadPlusWhole,
	"noteheadRectangularClusterBlackBottom": NoteheadRectangularClusterBlackBottom,
	"noteheadRectangularClusterBlackMiddle": NoteheadRectangularClusterBlackMiddle,
	"noteheadRectangularClusterBlackTop": NoteheadRectangularClusterBlackTop,
	"noteheadRectangularClusterWhiteBottom": NoteheadRectangularClusterWhiteBottom,
	"noteheadRectangularClusterWhiteMiddle": NoteheadRectangularClusterWhiteMiddle,
	"noteheadRectangularClusterWhiteTop": NoteheadRectangularClusterWhiteTop,
	"noteheadRoundBlack": NoteheadRoundBlack,
	"noteheadRoundBlackDoubleSlashed": NoteheadRoundBlackDoubleSlashed,
	"noteheadRoundBlackLarge": NoteheadRoundBlackLarge,
	"noteheadRoundBlackSlashed": NoteheadRoundBlackSlashed,
	"noteheadRoundBlackSlashedLarge": NoteheadRoundBlackSlashedLarge,
	"noteheadRoundWhite": NoteheadRoundWhite,
	"noteheadRoundWhiteDoubleSlashed": NoteheadRoundWhiteDoubleSlashed,
	"noteheadRoundWhiteLarge": NoteheadRoundWhiteLarge,
	"noteheadRoundWhiteSlashed": NoteheadRoundWhiteSlashed,
	"noteheadRoundWhiteSlashedLarge": NoteheadRoundWhiteSlashedLarge,
	"noteheadRoundWhiteWithDot": NoteheadRoundWhiteWithDot,
	"noteheadRoundWhiteWithDotLarge": NoteheadRoundWhiteWithDotLarge,
	"noteheadSlashDiamondWhite": NoteheadSlashDiamondWhite,
	"noteheadSlashHorizontalEnds": NoteheadSlashHorizontalEnds,
	"noteheadSlashHorizontalEndsMuted": NoteheadSlashHorizontalEndsMuted,
	"noteheadSlashVerticalEnds": NoteheadSlashVerticalEnds,
	"noteheadSlashVerticalEndsMuted": NoteheadSlashVerticalEndsMuted,
	"noteheadSlashVerticalEndsSmall": NoteheadSlashVerticalEndsSmall,
	"noteheadSlashWhiteDoubleWhole": NoteheadSlashWhiteDoubleWhole,
	"noteheadSlashWhiteHalf": NoteheadSlashWhiteHalf,
	"noteheadSlashWhiteMuted": NoteheadSlashWhiteMuted,
	"noteheadSlashWhiteWhole": NoteheadSlashWhiteWhole,
	"noteheadSlashX": NoteheadSlashX,
	"noteheadSlashedBlack1": NoteheadSlashedBlack1,
	"noteheadSlashedBlack2": NoteheadSlashedBlack2,
	"noteheadSlashedDoubleWhole1": NoteheadSlashedDoubleWhole1,
	"noteheadSlashedDoubleWhole2": NoteheadSlashedDoubleWhole2,
	"noteheadSlashedHalf1": NoteheadSlashedHalf1,
	"noteheadSlashedHalf2": NoteheadSlashedHalf2,
	"noteheadSlashedWhole1": NoteheadSlashedWhole1,
	"noteheadSlashedWhole2": NoteheadSlashedWhole2,
	"noteheadSquareBlack": NoteheadSquareBlack,
	"noteheadSquareBlackLarge": NoteheadSquareBlackLarge,
	"noteheadSquareBlackWhite": NoteheadSquareBlackWhite,
	"noteheadSquareWhite": NoteheadSquareWhite,
	"noteheadTriangleDownBlack": NoteheadTriangleDownBlack,
	"noteheadTriangleDownDoubleWhole": NoteheadTriangleDownDoubleWhole,
	"noteheadTriangleDownHalf": NoteheadTriangleDownHalf,
	"noteheadTriangleDownWhite": NoteheadTriangleDownWhite,
	"noteheadTriangleDownWhole": NoteheadTriangleDownWhole,
	"noteheadTriangleLeftBlack": NoteheadTriangleLeftBlack,
	"noteheadTriangleLeftWhite": NoteheadTriangleLeftWhite,
	"noteheadTriangleRightBlack": NoteheadTriangleRightBlack,
	"noteheadTriangleRightWhite": NoteheadTriangleRightWhite,
	"noteheadTriangleRoundDownBlack": NoteheadTriangleRoundDownBlack,
	"noteheadTriangleRoundDownWhite": NoteheadTriangleRoundDownWhite,
	"noteheadTriangleUpBlack": NoteheadTriangleUpBlack,
	"noteheadTriangleUpDoubleWhole": NoteheadTriangleUpDoubleWhole,
	"noteheadTriangleUpHalf": NoteheadTriangleUpHalf,
	"noteheadTriangleUpRightBlack": NoteheadTriangleUpRightBlack,
	"noteheadTriangleUpRightWhite": NoteheadTriangleUpRightWhite,
	"noteheadTriangleUpWhite": NoteheadTriangleUpWhite,
	"noteheadTriangleUpWhole": NoteheadTriangleUpWhole,
	"noteheadVoidWithX": NoteheadVoidWithX,
	"noteheadWhole": NoteheadWhole,
	"noteheadWholeFilled": NoteheadWholeFilled,
	"noteheadWholeWithX": NoteheadWholeWithX,
	"noteheadXBlack": NoteheadXBlack,
	"noteheadXDoubleWhole": NoteheadXDoubleWhole,
	"noteheadXHalf": NoteheadXHalf,
	"noteheadXOrnate": NoteheadXOrnate,
	"noteheadXOrnateEllipse": NoteheadXOrnateEllipse,
	"noteheadXWhole": NoteheadXWhole,
	"octaveBaselineA": OctaveBaselineA,
	"octaveBaselineB": OctaveBaselineB,
	"octaveBaselineM": OctaveBaselineM,
	"octaveBaselineV": OctaveBaselineV,
	"octaveBassa": OctaveBassa,
	"octaveLoco": OctaveLoco,
	"octaveParensLeft": OctaveParensLeft,
	"octaveParensRight": OctaveParensRight,
	"octaveSuperscriptA": OctaveSuperscriptA,
	"octaveSuperscriptB": OctaveSuperscriptB,
	"octaveSuperscriptM": OctaveSuperscriptM,
	"octaveSuperscriptV": OctaveSuperscriptV,
	"oneHandedRollStevens": OneHandedRollStevens,
	"organGerman2Fusae": OrganGerman2Fusae,
	"organGerman2Minimae": OrganGerman2Minimae,
	"organGerman2OctaveUp": OrganGerman2OctaveUp,
	"organGerman2Semifusae": OrganGerman2Semifusae,
	"organGerman2Semiminimae": OrganGerman2Semiminimae,
	"organGerman3Fusae": OrganGerman3Fusae,
	"organGerman3Minimae": OrganGerman3Minimae,
	"organGerman3Semifusae": OrganGerman3Semifusae,
	"organGerman3Semiminimae": OrganGerman3Semiminimae,
	"organGerman4Fusae": OrganGerman4Fusae,
	"organGerman4Minimae": OrganGerman4Minimae,
	"organGerman4Semifusae": OrganGerman4Semifusae,
	"organGerman4Semiminimae": OrganGerman4Semiminimae,
	"organGerman5Fusae": OrganGerman5Fusae,
	"organGerman5Minimae": OrganGerman5Minimae,
	"organGerman5Semifusae": OrganGerman5Semifusae,
	"organGerman5Semiminimae": OrganGerman5Semiminimae,
	"organGerman6Fusae": OrganGerman6Fusae,
	"organGerman6Minimae": OrganGerman6Minimae,
	"organGerman6Semifusae": OrganGerman6Semifusae,
	"organGerman6Semiminimae": OrganGerman6Semiminimae,
	"organGermanALower": OrganGermanALower,
	"organGermanAUpper": OrganGermanAUpper,
	"organGermanAugmentationDot": OrganGermanAugmentationDot,
	"organGermanBLower": OrganGermanBLower,
	"organGermanBUpper": OrganGermanBUpper,
	"organGermanBuxheimerBrevis2": OrganGermanBuxheimerBrevis2,
	"organGermanBuxheimerBrevis3": OrganGermanBuxheimerBrevis3,
	"organGermanBuxheimerMinimaRest": OrganGermanBuxheimerMinimaRest,
	"organGermanBuxheimerSemibrevis": OrganGermanBuxheimerSemibrevis,
	"organGermanBuxheimerSemibrevisRest": OrganGermanBuxheimerSemibrevisRest,
	"organGermanCLower": OrganGermanCLower,
	"organGermanCUpper": OrganGermanCUpper,
	"organGermanCisLower": OrganGermanCisLower,
	"organGermanCisUpper": OrganGermanCisUpper,
	"organGermanDLower": OrganGermanDLower,
	"organGermanDUpper": OrganGermanDUpper,
	"organGermanDisLower": OrganGermanDisLower,
	"organGermanDisUpper": OrganGermanDisUpper,
	"organGermanELower": OrganGermanELower,
	"organGermanEUpper": OrganGermanEUpper,
	"organGermanFLower": OrganGermanFLower,
	"organGermanFUpper": OrganGermanFUpper,
	"organGermanFisLower": OrganGermanFisLower,
	"organGermanFisUpper": OrganGermanFisUpper,
	"organGermanFusa": OrganGermanFusa,
	"organGermanFusaRest": OrganGermanFusaRest,
	"organGermanGLower": OrganGermanGLower,
	"organGermanGUpper": OrganGermanGUpper,
	"organGermanGisLower": OrganGermanGisLower,
	"organGermanGisUpper": OrganGermanGisUpper,
	"organGermanHLower": OrganGermanHLower,
	"organGermanHUpper": OrganGermanHUpper,
	"organGermanMinima": OrganGermanMinima,
	"organGermanMinimaRest": OrganGermanMinimaRest,
	"organGermanOctaveDown": OrganGermanOctaveDown,
	"organGermanOctaveUp": OrganGermanOctaveUp,
	"organGermanSemibrevis": OrganGermanSemibrevis,
	"organGermanSemibrevisRest": OrganGermanSemibrevisRest,
	"organGermanSemifusa": OrganGermanSemifusa,
	"organGermanSemifusaRest": OrganGermanSemifusaRest,
	"organGermanSemiminima": OrganGermanSemiminima,
	"organGermanSemiminimaRest": OrganGermanSemiminimaRest,
	"organGermanTie": OrganGermanTie,
	"ornamentBottomLeftConcaveStroke": OrnamentBottomLeftConcaveStroke,
	"ornamentBottomLeftConcaveStrokeLarge": OrnamentBottomLeftConcaveStrokeLarge,
	"ornamentBottomLeftConvexStroke": OrnamentBottomLeftConvexStroke,
	"ornamentBottomRightConcaveStroke": OrnamentBottomRightConcaveStroke,
	"ornamentBottomRightConvexStroke": OrnamentBottomRightConvexStroke,
	"ornamentComma": OrnamentComma,
	"ornamentDoubleObliqueLinesAfterNote": OrnamentDoubleObliqueLinesAfterNote,
	"ornamentDoubleObliqueLinesBeforeNote": OrnamentDoubleObliqueLinesBeforeNote,
	"ornamentDownCurve": OrnamentDownCurve,
	"ornamentHaydn": OrnamentHaydn,
	"ornamentHighLeftConcaveStroke": OrnamentHighLeftConcaveStroke,
	"ornamentHighLeftConvexStroke": OrnamentHighLeftConvexStroke,
	"ornamentHighRightConcaveStroke": OrnamentHighRightConcaveStroke,
	"ornamentHighRightConvexStroke": OrnamentHighRightConvexStroke,
	"ornamentHookAfterNote": OrnamentHookAfterNote,
	"ornamentHookBeforeNote": OrnamentHookBeforeNote,
	"ornamentLeftFacingHalfCircle": OrnamentLeftFacingHalfCircle,
	"ornamentLeftFacingHook": OrnamentLeftFacingHook,
	"ornamentLeftPlus": OrnamentLeftPlus,
	"ornamentLeftShakeT": OrnamentLeftShakeT,
	"ornamentLeftVerticalStroke": OrnamentLeftVerticalStroke,
	"ornamentLeftVerticalStrokeWithCross": OrnamentLeftVerticalStrokeWithCross,
	"ornamentLowLeftConcaveStroke": OrnamentLowLeftConcaveStroke,
	"ornamentLowLeftConvexStroke": OrnamentLowLeftConvexStroke,
	"ornamentLowRightConcaveStroke": OrnamentLowRightConcaveStroke,
	"ornamentLowRightConvexStroke": OrnamentLowRightConvexStroke,
	"ornamentMiddleVerticalStroke": OrnamentMiddleVerticalStroke,
	"ornamentMordent": OrnamentMordent,
	"ornamentObliqueLineAfterNote": OrnamentObliqueLineAfterNote,
	"ornamentObliqueLineBeforeNote": OrnamentObliqueLineBeforeNote,
	"ornamentObliqueLineHorizAfterNote": OrnamentObliqueLineHorizAfterNote,
	"ornamentObliqueLineHorizBeforeNote": OrnamentObliqueLineHorizBeforeNote,
	"ornamentOriscus": OrnamentOriscus,
	"ornamentPinceCouperin": OrnamentPinceCouperin,
	"ornamentPortDeVoixV": OrnamentPortDeVoixV,
	"ornamentPrecompAppoggTrill": OrnamentPrecompAppoggTrill,
	"ornamentPrecompAppoggTrillSuffix": OrnamentPrecompAppoggTrillSuffix,
	"ornamentPrecompCadence": OrnamentPrecompCadence,
	"ornamentPrecompCadenceUpperPrefix": OrnamentPrecompCadenceUpperPrefix,
	"ornamentPrecompCadenceUpperPrefixTurn": OrnamentPrecompCadenceUpperPrefixTurn,
	"ornamentPrecompCadenceWithTurn": OrnamentPrecompCadenceWithTurn,
	"ornamentPrecompDescendingSlide": OrnamentPrecompDescendingSlide,
	"ornamentPrecompDoubleCadenceLowerPrefix": OrnamentPrecompDoubleCadenceLowerPrefix,
	"ornamentPrecompDoubleCadenceUpperPrefix": OrnamentPrecompDoubleCadenceUpperPrefix,
	"ornamentPrecompDoubleCadenceUpperPrefixTurn": OrnamentPrecompDoubleCadenceUpperPrefixTurn,
	"ornamentPrecompInvertedMordentUpperPrefix": OrnamentPrecompInvertedMordentUpperPrefix,
	"ornamentPrecompMordentRelease": OrnamentPrecompMordentRelease,
	"ornamentPrecompMordentUpperPrefix": OrnamentPrecompMordentUpperPrefix,
	"ornamentPrecompPortDeVoixMordent": OrnamentPrecompPortDeVoixMordent,
	"ornamentPrecompSlide": OrnamentPrecompSlide,
	"ornamentPrecompSlideTrillBach": OrnamentPrecompSlideTrillBach,
	"ornamentPrecompSlideTrillDAnglebert": OrnamentPrecompSlideTrillDAnglebert,
	"ornamentPrecompSlideTrillMarpurg": OrnamentPrecompSlideTrillMarpurg,
	"ornamentPrecompSlideTrillMuffat": OrnamentPrecompSlideTrillMuffat,
	"ornamentPrecompSlideTrillSuffixMuffat": OrnamentPrecompSlideTrillSuffixMuffat,
	"ornamentPrecompTrillLowerSuffix": OrnamentPrecompTrillLowerSuffix,
	"ornamentPrecompTrillSuffixDandrieu": OrnamentPrecompTrillSuffixDandrieu,
	"ornamentPrecompTrillWithMordent": OrnamentPrecompTrillWithMordent,
	"ornamentPrecompTurnTrillBach": OrnamentPrecompTurnTrillBach,
	"ornamentPrecompTurnTrillDAnglebert": OrnamentPrecompTurnTrillDAnglebert,
	"ornamentQuilisma": OrnamentQuilisma,
	"ornamentRightFacingHalfCircle": OrnamentRightFacingHalfCircle,
	"ornamentRightFacingHook": OrnamentRightFacingHook,
	"ornamentRightVerticalStroke": OrnamentRightVerticalStroke,
	"ornamentSchleifer": OrnamentSchleifer,
	"ornamentShake3": OrnamentShake3,
	"ornamentShakeMuffat1": OrnamentShakeMuffat1,
	"ornamentShortObliqueLineAfterNote": OrnamentShortObliqueLineAfterNote,
	"ornamentShortObliqueLineBeforeNote": OrnamentShortObliqueLineBeforeNote,
	"ornamentShortTrill": OrnamentShortTrill,
	"ornamentTopLeftConcaveStroke": OrnamentTopLeftConcaveStroke,
	"ornamentTopLeftConvexStroke": OrnamentTopLeftConvexStroke,
	"ornamentTopRightConcaveStroke": OrnamentTopRightConcaveStroke,
	"ornamentTopRightConvexStroke": OrnamentTopRightConvexStroke,
	"ornamentTremblement": OrnamentTremblement,
	"ornamentTremblementCouperin": OrnamentTremblementCouperin,
	"ornamentTrill": OrnamentTrill,
	"ornamentTurn": OrnamentTurn,
	"ornamentTurnInverted": OrnamentTurnInverted,
	"ornamentTurnSlash": OrnamentTurnSlash,
	"ornamentTurnUp": OrnamentTurnUp,
	"ornamentTurnUpS": OrnamentTurnUpS,
	"ornamentUpCurve": OrnamentUpCurve,
	"ornamentVerticalLine": OrnamentVerticalLine,
	"ornamentZigZagLineNoRightEnd": OrnamentZigZagLineNoRightEnd,
	"ornamentZigZagLineWithRightEnd": OrnamentZigZagLineWithRightEnd,
	"ottava": Ottava,
	"ottavaAlta": OttavaAlta,
	"ottavaBassa": OttavaBassa,
	"ottavaBassaBa": OttavaBassaBa,
	"ottavaBassaVb": OttavaBassaVb,
	"pendereckiTremolo": PendereckiTremolo,
	"pictAgogo": PictAgogo,
	"pictAlmglocken": PictAlmglocken,
	"pictAnvil": PictAnvil,
	"pictBambooChimes": PictBambooChimes,
	"pictBambooScraper": PictBambooScraper,
	"pictBassDrum": PictBassDrum,
	"pictBassDrumOnSide": PictBassDrumOnSide,
	"pictBeaterBow": PictBeaterBow,
	"pictBeaterBox": PictBeaterBox,
	"pictBeaterBrassMalletsDown": PictBeaterBrassMalletsDown,
	"pictBeaterBrassMalletsLeft": PictBeaterBrassMalletsLeft,
	"pictBeaterBrassMalletsRight": PictBeaterBrassMalletsRight,
	"pictBeaterBrassMalletsUp": PictBeaterBrassMalletsUp,
	"pictBeaterCombiningDashedCircle": PictBeaterCombiningDashedCircle,
	"pictBeaterCombiningParentheses": PictBeaterCombiningParentheses,
	"pictBeaterDoubleBassDrumDown": PictBeaterDoubleBassDrumDown,
	"pictBeaterDoubleBassDrumUp": PictBeaterDoubleBassDrumUp,
	"pictBeaterFinger": PictBeaterFinger,
	"pictBeaterFingernails": PictBeaterFingernails,
	"pictBeaterFist": PictBeaterFist,
	"pictBeaterGuiroScraper": PictBeaterGuiroScraper,
	"pictBeaterHammer": PictBeaterHammer,
	"pictBeaterHammerMetalDown": PictBeaterHammerMetalDown,
	"pictBeaterHammerMetalUp": PictBeaterHammerMetalUp,
	"pictBeaterHammerPlasticDown": PictBeaterHammerPlasticDown,
	"pictBeaterHammerPlasticUp": PictBeaterHammerPlasticUp,
	"pictBeaterHammerWoodDown": PictBeaterHammerWoodDown,
	"pictBeaterHammerWoodUp": PictBeaterHammerWoodUp,
	"pictBeaterHand": PictBeaterHand,
	"pictBeaterHardBassDrumDown": PictBeaterHardBassDrumDown,
	"pictBeaterHardBassDrumUp": PictBeaterHardBassDrumUp,
	"pictBeaterHardGlockenspielDown": PictBeaterHardGlockenspielDown,
	"pictBeaterHardGlockenspielLeft": PictBeaterHardGlockenspielLeft,
	"pictBeaterHardGlockenspielRight": PictBeaterHardGlockenspielRight,
	"pictBeaterHardGlockenspielUp": PictBeaterHardGlockenspielUp,
	"pictBeaterHardTimpaniDown": PictBeaterHardTimpaniDown,
	"pictBeaterHardTimpaniLeft": PictBeaterHardTimpaniLeft,
	"pictBeaterHardTimpaniRight": PictBeaterHardTimpaniRight,
	"pictBeaterHardTimpaniUp": PictBeaterHardTimpaniUp,
	"pictBeaterHardXylophoneDown": PictBeaterHardXylophoneDown,
	"pictBeaterHardXylophoneLeft": PictBeaterHardXylophoneLeft,
	"pictBeaterHardXylophoneRight": PictBeaterHardXylophoneRight,
	"pictBeaterHardXylophoneUp": PictBeaterHardXylophoneUp,
	"pictBeaterHardYarnDown": PictBeaterHardYarnDown,
	"pictBeaterHardYarnLeft": PictBeaterHardYarnLeft,
	"pictBeaterHardYarnRight": PictBeaterHardYarnRight,
	"pictBeaterHardYarnUp": PictBeaterHardYarnUp,
	"pictBeaterJazzSticksDown": PictBeaterJazzSticksDown,
	"pictBeaterJazzSticksUp": PictBeaterJazzSticksUp,
	"pictBeaterKnittingNeedle": PictBeaterKnittingNeedle,
	"pictBeaterMallet": PictBeaterMallet,
	"pictBeaterMalletDown": PictBeaterMalletDown,
	"pictBeaterMediumBassDrumDown": PictBeaterMediumBassDrumDown,
	"pictBeaterMediumBassDrumUp": PictBeaterMediumBassDrumUp,
	"pictBeaterMediumTimpaniDown": PictBeaterMediumTimpaniDown,
	"pictBeaterMediumTimpaniLeft": PictBeaterMediumTimpaniLeft,
	"pictBeaterMediumTimpaniRight": PictBeaterMediumTimpaniRight,
	"pictBeaterMediumTimpaniUp": PictBeaterMediumTimpaniUp,
	"pictBeaterMediumXylophoneDown": PictBeaterMediumXylophoneDown,
	"pictBeaterMediumXylophoneLeft": PictBeaterMediumXylophoneLeft,
	"pictBeaterMediumXylophoneRight": PictBeaterMediumXylophoneRight,
	"pictBeaterMediumXylophoneUp": PictBeaterMediumXylophoneUp,
	"pictBeaterMediumYarnDown": PictBeaterMediumYarnDown,
	"pictBeaterMediumYarnLeft": PictBeaterMediumYarnLeft,
	"pictBeaterMediumYarnRight": PictBeaterMediumYarnRight,
	"pictBeaterMediumYarnUp": PictBeaterMediumYarnUp,
	"pictBeaterMetalBassDrumDown": PictBeaterMetalBassDrumDown,
	"pictBeaterMetalBassDrumUp": PictBeaterMetalBassDrumUp,
	"pictBeaterMetalDown": PictBeaterMetalDown,
	"pictBeaterMetalHammer": PictBeaterMetalHammer,
	"pictBeaterMetalLeft": PictBeaterMetalLeft,
	"pictBeaterMetalRight": PictBeaterMetalRight,
	"pictBeaterMetalUp": PictBeaterMetalUp,
	"pictBeaterSnareSticksDown": PictBeaterSnareSticksDown,
	"pictBeaterSnareSticksUp": PictBeaterSnareSticksUp,
	"pictBeaterSoftBassDrumDown": PictBeaterSoftBassDrumDown,
	"pictBeaterSoftBassDrumUp": PictBeaterSoftBassDrumUp,
	"pictBeaterSoftGlockenspielDown": PictBeaterSoftGlockenspielDown,
	"pictBeaterSoftGlockenspielLeft": PictBeaterSoftGlockenspielLeft,
	"pictBeaterSoftGlockenspielRight": PictBeaterSoftGlockenspielRight,
	"pictBeaterSoftGlockenspielUp": PictBeaterSoftGlockenspielUp,
	"pictBeaterSoftTimpaniDown": PictBeaterSoftTimpaniDown,
	"pictBeaterSoftTimpaniLeft": PictBeaterSoftTimpaniLeft,
	"pictBeaterSoftTimpaniRight": PictBeaterSoftTimpaniRight,
	"pictBeaterSoftTimpaniUp": PictBeaterSoftTimpaniUp,
	"pictBeaterSoftXylophone": PictBeaterSoftXylophone,
	"pictBeaterSoftXylophoneDown": PictBeaterSoftXylophoneDown,
	"pictBeaterSoftXylophoneLeft": PictBeaterSoftXylophoneLeft,
	"pictBeaterSoftXylophoneRight": PictBeaterSoftXylophoneRight,
	"pictBeaterSoftXylophoneUp": PictBeaterSoftXylophoneUp,
	"pictBeaterSoftYarnDown": PictBeaterSoftYarnDown,
	"pictBeaterSoftYarnLeft": PictBeaterSoftYarnLeft,
	"pictBeaterSoftYarnRight": PictBeaterSoftYarnRight,
	"pictBeaterSoftYarnUp": PictBeaterSoftYarnUp,
	"pictBeaterSpoonWoodenMallet": PictBeaterSpoonWoodenMallet,
	"pictBeaterSuperballDown": PictBeaterSuperballDown,
	"pictBeaterSuperballLeft": PictBeaterSuperballLeft,
	"pictBeaterSuperballRight": PictBeaterSuperballRight,
	"pictBeaterSuperballUp": PictBeaterSuperballUp,
	"pictBeaterTriangleDown": PictBeaterTriangleDown,
	"pictBeaterTrianglePlain": PictBeaterTrianglePlain,
	"pictBeaterTriangleUp": PictBeaterTriangleUp,
	"pictBeaterWireBrushesDown": PictBeaterWireBrushesDown,
	"pictBeaterWireBrushesUp": PictBeaterWireBrushesUp,
	"pictBeaterWoodTimpaniDown": PictBeaterWoodTimpaniDown,
	"pictBeaterWoodTimpaniLeft": PictBeaterWoodTimpaniLeft,
	"pictBeaterWoodTimpaniRight": PictBeaterWoodTimpaniRight,
	"pictBeaterWoodTimpaniUp": PictBeaterWoodTimpaniUp,
	"pictBeaterWoodXylophoneDown": PictBeaterWoodXylophoneDown,
	"pictBeaterWoodXylophoneLeft": PictBeaterWoodXylophoneLeft,
	"pictBeaterWoodXylophoneRight": PictBeaterWoodXylophoneRight,
	"pictBeaterWoodXylophoneUp": PictBeaterWoodXylophoneUp,
	"pictBell": PictBell,
	"pictBellOfCymbal": PictBellOfCymbal,
	"pictBellPlate": PictBellPlate,
	"pictBellTree": PictBellTree,
	"pictBirdWhistle": PictBirdWhistle,
	"pictBoardClapper": PictBoardClapper,
	"pictBongos": PictBongos,
	"pictBrakeDrum": PictBrakeDrum,
	"pictCabasa": PictCabasa,
	"pictCannon": PictCannon,
	"pictCarHorn": PictCarHorn,
	"pictCastanets": PictCastanets,
	"pictCastanetsWithHandle": PictCastanetsWithHandle,
	"pictCelesta": PictCelesta,
	"pictCencerro": PictCencerro,
	"pictCenter1": PictCenter1,
	"pictCenter2": PictCenter2,
	"pictCenter3": PictCenter3,
	"pictChainRattle": PictChainRattle,
	"pictChimes": PictChimes,
	"pictChineseCymbal": PictChineseCymbal,
	"pictChokeCymbal": PictChokeCymbal,
	"pictClaves": PictClaves,
	"pictCoins": PictCoins,
	"pictConga": PictConga,
	"pictCowBell": PictCowBell,
	"pictCrashCymbals": PictCrashCymbals,
	"pictCrotales": PictCrotales,
	"pictCrushStem": PictCrushStem,
	"pictCuica": PictCuica,
	"pictCymbalTongs": PictCymbalTongs,
	"pictDamp1": PictDamp1,
	"pictDamp2": PictDamp2,
	"pictDamp3": PictDamp3,
	"pictDamp4": PictDamp4,
	"pictDeadNoteStem": PictDeadNoteStem,
	"pictDrumStick": PictDrumStick,
	"pictDuckCall": PictDuckCall,
	"pictEdgeOfCymbal": PictEdgeOfCymbal,
	"pictEmptyTrap": PictEmptyTrap,
	"pictFingerCymbals": PictFingerCymbals,
	"pictFlexatone": PictFlexatone,
	"pictFootballRatchet": PictFootballRatchet,
	"pictGlassHarmonica": PictGlassHarmonica,
	"pictGlassHarp": PictGlassHarp,
	"pictGlassPlateChimes": PictGlassPlateChimes,
	"pictGlassTubeChimes": PictGlassTubeChimes,
	"pictGlsp": PictGlsp,
	"pictGlspSmithBrindle": PictGlspSmithBrindle,
	"pictGobletDrum": PictGobletDrum,
	"pictGong": PictGong,
	"pictGongWithButton": PictGongWithButton,
	"pictGuiro": PictGuiro,
	"pictGumHardDown": PictGumHardDown,
	"pictGumHardLeft": PictGumHardLeft,
	"pictGumHardRight": PictGumHardRight,
	"pictGumHardUp": PictGumHardUp,
	"pictGumMediumDown": PictGumMediumDown,
	"pictGumMediumLeft": PictGumMediumLeft,
	"pictGumMediumRight": PictGumMediumRight,
	"pictGumMediumUp": PictGumMediumUp,
	"pictGumSoftDown": PictGumSoftDown,
	"pictGumSoftLeft": PictGumSoftLeft,
	"pictGumSoftRight": PictGumSoftRight,
	"pictGumSoftUp": PictGumSoftUp,
	"pictHalfOpen1": PictHalfOpen1,
	"pictHalfOpen2": PictHalfOpen2,
	"pictHandbell": PictHandbell,
	"pictHiHat": PictHiHat,
	"pictHiHatOnStand": PictHiHatOnStand,
	"pictJawHarp": PictJawHarp,
	"pictJingleBells": PictJingleBells,
	"pictKlaxonHorn": PictKlaxonHorn,
	"pictLeftHandCircle": PictLeftHandCircle,
	"pictLionsRoar": PictLionsRoar,
	"pictLithophone": PictLithophone,
	"pictLogDrum": PictLogDrum,
	"pictLotusFlute": PictLotusFlute,
	"pictMar": PictMar,
	"pictMarSmithBrindle": PictMarSmithBrindle,
	"pictMaraca": PictMaraca,
	"pictMaracas": PictMaracas,
	"pictMegaphone": PictMegaphone,
	"pictMetalPlateChimes": PictMetalPlateChimes,
	"pictMetalTubeChimes": PictMetalTubeChimes,
	"pictMusicalSaw": PictMusicalSaw,
	"pictNormalPosition": PictNormalPosition,
	"pictOnRim": PictOnRim,
	"pictOpen": PictOpen,
	"pictOpenRimShot": PictOpenRimShot,
	"pictPistolShot": PictPistolShot,
	"pictPoliceWhistle": PictPoliceWhistle,
	"pictQuijada": PictQuijada,
	"pictRainstick": PictRainstick,
	"pictRatchet": PictRatchet,
	"pictRecoReco": PictRecoReco,
	"pictRightHandSquare": PictRightHandSquare,
	"pictRim1": PictRim1,
	"pictRim2": PictRim2,
	"pictRim3": PictRim3,
	"pictRimShotOnStem": PictRimShotOnStem,
	"pictSandpaperBlocks": PictSandpaperBlocks,
	"pictScrapeAroundRim": PictScrapeAroundRim,
	"pictScrapeAroundRimClockwise": PictScrapeAroundRimClockwise,
	"pictScrapeCenterToEdge": PictScrapeCenterToEdge,
	"pictScrapeEdgeToCenter": PictScrapeEdgeToCenter,
	"pictShellBells": PictShellBells,
	"pictShellChimes": PictShellChimes,
	"pictSiren": PictSiren,
	"pictSistrum": PictSistrum,
	"pictSizzleCymbal": PictSizzleCymbal,
	"pictSleighBell": PictSleighBell,
	"pictSlideBrushOnGong": PictSlideBrushOnGong,
	"pictSlideWhistle": PictSlideWhistle,
	"pictSlitDrum": PictSlitDrum,
	"pictSnareDrum": PictSnareDrum,
	"pictSnareDrumMilitary": PictSnareDrumMilitary,
	"pictSnareDrumSnaresOff": PictSnareDrumSnaresOff,
	"pictSteelDrums": PictSteelDrums,
	"pictStickShot": PictStickShot,
	"pictSuperball": PictSuperball,
	"pictSuspendedCymbal": PictSuspendedCymbal,
	"pictSwishStem": PictSwishStem,
	"pictTabla": PictTabla,
	"pictTamTam": PictTamTam,
	"pictTamTamWithBeater": PictTamTamWithBeater,
	"pictTambourine": PictTambourine,
	"pictTempleBlocks": PictTempleBlocks,
	"pictTenorDrum": PictTenorDrum,
	"pictThundersheet": PictThundersheet,
	"pictTimbales": PictTimbales,
	"pictTimpani": PictTimpani,
	"pictTomTom": PictTomTom,
	"pictTomTomChinese": PictTomTomChinese,
	"pictTomTomIndoAmerican": PictTomTomIndoAmerican,
	"pictTomTomJapanese": PictTomTomJapanese,
	"pictTriangle": PictTriangle,
	"pictTubaphone": PictTubaphone,
	"pictTubularBells": PictTubularBells,
	"pictTurnLeftStem": PictTurnLeftStem,
	"pictTurnRightLeftStem": PictTurnRightLeftStem,
	"pictTurnRightStem": PictTurnRightStem,
	"pictVib": PictVib,
	"pictVibMotorOff": PictVibMotorOff,
	"pictVibSmithBrindle": PictVibSmithBrindle,
	"pictVibraslap": PictVibraslap,
	"pictVietnameseHat": PictVietnameseHat,
	"pictWhip": PictWhip,
	"pictWindChimesGlass": PictWindChimesGlass,
	"pictWindMachine": PictWindMachine,
	"pictWindWhistle": PictWindWhistle,
	"pictWoodBlock": PictWoodBlock,
	"pictWoundHardDown": PictWoundHardDown,
	"pictWoundHardLeft": PictWoundHardLeft,
	"pictWoundHardRight": PictWoundHardRight,
	"pictWoundHardUp": PictWoundHardUp,
	"pictWoundSoftDown": PictWoundSoftDown,
	"pictWoundSoftLeft": PictWoundSoftLeft,
	"pictWoundSoftRight": PictWoundSoftRight,
	"pictWoundSoftUp": PictWoundSoftUp,
	"pictXyl": PictXyl,
	"pictXylBass": PictXylBass,
	"pictXylSmithBrindle": PictXylSmithBrindle,
	"pictXylTenor": PictXylTenor,
	"pictXylTenorTrough": PictXylTenorTrough,
	"pictXylTrough": PictXylTrough,
	"pluckedBuzzPizzicato": PluckedBuzzPizzicato,
	"pluckedDamp": PluckedDamp,
	"pluckedDampAll": PluckedDampAll,
	"pluckedDampOnStem": PluckedDampOnStem,
	"pluckedFingernailFlick": PluckedFingernailFlick,
	"pluckedLeftHandPizzicato": PluckedLeftHandPizzicato,
	"pluckedPlectrum": PluckedPlectrum,
	"pluckedSnapPizzicatoAbove": PluckedSnapPizzicatoAbove,
	"pluckedSnapPizzicatoBelow": PluckedSnapPizzicatoBelow,
	"pluckedWithFingernails": PluckedWithFingernails,
	"quindicesima": Quindicesima,
	"quindicesimaAlta": QuindicesimaAlta,
	"quindicesimaBassa": QuindicesimaBassa,
	"quindicesimaBassaMb": QuindicesimaBassaMb,
	"repeat1Bar": Repeat1Bar,
	"repeat2Bars": Repeat2Bars,
	"repeat4Bars": Repeat4Bars,
	"repeatBarLowerDot": RepeatBarLowerDot,
	"repeatBarSlash": RepeatBarSlash,
	"repeatBarUpperDot": RepeatBarUpperDot,
	"repeatDot": RepeatDot,
	"repeatDots": RepeatDots,
	"repeatLeft": RepeatLeft,
	"repeatRight": RepeatRight,
	"repeatRightLeft": RepeatRightLeft,
	"rest1024th": Rest1024th,
	"rest128th": Rest128th,
	"rest16th": Rest16th,
	"rest256th": Rest256th,
	"rest32nd": Rest32nd,
	"rest512th": Rest512th,
	"rest64th": Rest64th,
	"rest8th": Rest8th,
	"restDoubleWhole": RestDoubleWhole,
	"restDoubleWholeLegerLine": RestDoubleWholeLegerLine,
	"restHBar": RestHBar,
	"restHBarLeft": RestHBarLeft,
	"restHBarMiddle": RestHBarMiddle,
	"restHBarRight": RestHBarRight,
	"restHalf": RestHalf,
	"restHalfLegerLine": RestHalfLegerLine,
	"restLonga": RestLonga,
	"restMaxima": RestMaxima,
	"restQuarter": RestQuarter,
	"restQuarterOld": RestQuarterOld,
	"restQuarterZ": RestQuarterZ,
	"restWhole": RestWhole,
	"restWholeLegerLine": RestWholeLegerLine,
	"reversedBrace": ReversedBrace,
	"reversedBracketBottom": ReversedBracketBottom,
	"reversedBracketTop": ReversedBracketTop,
	"rightRepeatSmall": RightRepeatSmall,
	"scaleDegree1": ScaleDegree1,
	"scaleDegree2": ScaleDegree2,
	"scaleDegree3": ScaleDegree3,
	"scaleDegree4": ScaleDegree4,
	"scaleDegree5": ScaleDegree5,
	"scaleDegree6": ScaleDegree6,
	"scaleDegree7": ScaleDegree7,
	"scaleDegree8": ScaleDegree8,
	"scaleDegree9": ScaleDegree9,
	"schaefferClef": SchaefferClef,
	"schaefferFClefToGClef": SchaefferFClefToGClef,
	"schaefferGClefToFClef": SchaefferGClefToFClef,
	"schaefferPreviousClef": SchaefferPreviousClef,
	"segno": Segno,
	"segnoSerpent1": SegnoSerpent1,
	"segnoSerpent2": SegnoSerpent2,
	"semipitchedPercussionClef1": SemipitchedPercussionClef1,
	"semipitchedPercussionClef2": SemipitchedPercussionClef2,
	"smnFlat": SmnFlat,
	"smnFlatWhite": SmnFlatWhite,
	"smnHistoryDoubleFlat": SmnHistoryDoubleFlat,
	"smnHistoryDoubleSharp": SmnHistoryDoubleSharp,
	"smnHistoryFlat": SmnHistoryFlat,
	"smnHistorySharp": SmnHistorySharp,
	"smnNatural": SmnNatural,
	"smnSharp": SmnSharp,
	"smnSharpDown": SmnSharpDown,
	"smnSharpWhite": SmnSharpWhite,
	"smnSharpWhiteDown": SmnSharpWhiteDown,
	"splitBarDivider": SplitBarDivider,
	"staff1Line": Staff1Line,
	"staff1LineNarrow": Staff1LineNarrow,
	"staff1LineWide": Staff1LineWide,
	"staff2Lines": Staff2Lines,
	"staff2LinesNarrow": Staff2LinesNarrow,
	"staff2LinesWide": Staff2LinesWide,
	"staff3Lines": Staff3Lines,
	"staff3LinesNarrow": Staff3LinesNarrow,
	"staff3LinesWide": Staff3LinesWide,
	"staff4Lines": Staff4Lines,
	"staff4LinesNarrow": Staff4LinesNarrow,
	"staff4LinesWide": Staff4LinesWide,
	"staff5Lines": Staff5Lines,
	"staff5LinesNarrow": Staff5LinesNarrow,
	"staff5LinesWide": Staff5LinesWide,
	"staff6Lines": Staff6Lines,
	"staff6LinesNarrow": Staff6LinesNarrow,
	"staff6LinesWide": Staff6LinesWide,
	"staffDivideArrowDown": StaffDivideArrowDown,
	"staffDivideArrowUp": StaffDivideArrowUp,
	"staffDivideArrowUpDown": StaffDivideArrowUpDown,
	"staffPosLower1": StaffPosLower1,
	"staffPosLower2": StaffPosLower2,
	"staffPosLower3": StaffPosLower3,
	"staffPosLower4": StaffPosLower4,
	"staffPosLower5": StaffPosLower5,
	"staffPosLower6": StaffPosLower6,
	"staffPosLower7": StaffPosLower7,
	"staffPosLower8": StaffPosLower8,
	"staffPosRaise1": StaffPosRaise1,
	"staffPosRaise2": StaffPosRaise2,
	"staffPosRaise3": StaffPosRaise3,
	"staffPosRaise4": StaffPosRaise4,
	"staffPosRaise5": StaffPosRaise5,
	"staffPosRaise6": StaffPosRaise6,
	"staffPosRaise7": StaffPosRaise7,
	"staffPosRaise8": StaffPosRaise8,
	"stem": Stem,
	"stemBowOnBridge": StemBowOnBridge,
	"stemBowOnTailpiece": StemBowOnTailpiece,
	"stemBuzzRoll": StemBuzzRoll,
	"stemDamp": StemDamp,
	"stemHarpStringNoise": StemHarpStringNoise,
	"stemMultiphonicsBlack": StemMultiphonicsBlack,
	"stemMultiphonicsBlackWhite": StemMultiphonicsBlackWhite,
	"stemMultiphonicsWhite": StemMultiphonicsWhite,
	"stemPendereckiTremolo": StemPendereckiTremolo,
	"stemRimShot": StemRimShot,
	"stemSprechgesang": StemSprechgesang,
	"stemSulPonticello": StemSulPonticello,
	"stemSussurando": StemSussurando,
	"stemSwished": StemSwished,
	"stemVibratoPulse": StemVibratoPulse,
	"stockhausenTremolo": StockhausenTremolo,
	"stringsBowBehindBridge": StringsBowBehindBridge,
	"stringsBowBehindBridgeFourStrings": StringsBowBehindBridgeFourStrings,
	"stringsBowBehindBridgeOneString": StringsBowBehindBridgeOneString,
	"stringsBowBehindBridgeThreeStrings": StringsBowBehindBridgeThreeStrings,
	"stringsBowBehindBridgeTwoStrings": StringsBowBehindBridgeTwoStrings,
	"stringsBowOnBridge": StringsBowOnBridge,
	"stringsBowOnTailpiece": StringsBowOnTailpiece,
	"stringsChangeBowDirection": StringsChangeBowDirection,
	"stringsDownBow": StringsDownBow,
	"stringsDownBowAwayFromBody": StringsDownBowAwayFromBody,
	"stringsDownBowBeyondBridge": StringsDownBowBeyondBridge,
	"stringsDownBowTowardsBody": StringsDownBowTowardsBody,
	"stringsDownBowTurned": StringsDownBowTurned,
	"stringsFouette": StringsFouette,
	"stringsHalfHarmonic": StringsHalfHarmonic,
	"stringsHarmonic": StringsHarmonic,
	"stringsJeteAbove": StringsJeteAbove,
	"stringsJeteBelow": StringsJeteBelow,
	"stringsMuteOff": StringsMuteOff,
	"stringsMuteOn": StringsMuteOn,
	"stringsOverpressureDownBow": StringsOverpressureDownBow,
	"stringsOverpressureNoDirection": StringsOverpressureNoDirection,
	"stringsOverpressurePossibileDownBow": StringsOverpressurePossibileDownBow,
	"stringsOverpressurePossibileUpBow": StringsOverpressurePossibileUpBow,
	"stringsOverpressureUpBow": StringsOverpressureUpBow,
	"stringsScrapeCircularClockwise": StringsScrapeCircularClockwise,
	"stringsScrapeCircularCounterclockwise": StringsScrapeCircularCounterclockwise,
	"stringsScrapeParallelInward": StringsScrapeParallelInward,
	"stringsScrapeParallelOutward": StringsScrapeParallelOutward,
	"stringsThumbPosition": StringsThumbPosition,
	"stringsThumbPositionTurned": StringsThumbPositionTurned,
	"stringsTripleChopInward": StringsTripleChopInward,
	"stringsTripleChopOutward": StringsTripleChopOutward,
	"stringsUpBow": StringsUpBow,
	"stringsUpBowAwayFromBody": StringsUpBowAwayFromBody,
	"stringsUpBowBeyondBridge": StringsUpBowBeyondBridge,
	"stringsUpBowTowardsBody": StringsUpBowTowardsBody,
	"stringsUpBowTurned": StringsUpBowTurned,
	"stringsVibratoPulse": StringsVibratoPulse,
	"swissRudimentsNoteheadBlackDouble": SwissRudimentsNoteheadBlackDouble,
	"swissRudimentsNoteheadBlackFlam": SwissRudimentsNoteheadBlackFlam,
	"swissRudimentsNoteheadHalfDouble": SwissRudimentsNoteheadHalfDouble,
	"swissRudimentsNoteheadHalfFlam": SwissRudimentsNoteheadHalfFlam,
	"systemDivider": SystemDivider,
	"systemDividerExtraLong": SystemDividerExtraLong,
	"systemDividerLong": SystemDividerLong,
	"textAugmentationDot": TextAugmentationDot,
	"textBlackNoteFrac16thLongStem": TextBlackNoteFrac16thLongStem,
	"textBlackNoteFrac16thShortStem": TextBlackNoteFrac16thShortStem,
	"textBlackNoteFrac32ndLongStem": TextBlackNoteFrac32ndLongStem,
	"textBlackNoteFrac8thLongStem": TextBlackNoteFrac8thLongStem,
	"textBlackNoteFrac8thShortStem": TextBlackNoteFrac8thShortStem,
	"textBlackNoteLongStem": TextBlackNoteLongStem,
	"textBlackNoteShortStem": TextBlackNoteShortStem,
	"textCont16thBeamLongStem": TextCont16thBeamLongStem,
	"textCont16thBeamShortStem": TextCont16thBeamShortStem,
	"textCont32ndBeamLongStem": TextCont32ndBeamLongStem,
	"textCont8thBeamLongStem": TextCont8thBeamLongStem,
	"textCont8thBeamShortStem": TextCont8thBeamShortStem,
	"textHeadlessBlackNoteFrac16thLongStem": TextHeadlessBlackNoteFrac16thLongStem,
	"textHeadlessBlackNoteFrac16thShortStem": TextHeadlessBlackNoteFrac16thShortStem,
	"textHeadlessBlackNoteFrac32ndLongStem": TextHeadlessBlackNoteFrac32ndLongStem,
	"textHeadlessBlackNoteFrac8thLongStem": TextHeadlessBlackNoteFrac8thLongStem,
	"textHeadlessBlackNoteFrac8thShortStem": TextHeadlessBlackNoteFrac8thShortStem,
	"textHeadlessBlackNoteLongStem": TextHeadlessBlackNoteLongStem,
	"textHeadlessBlackNoteShortStem": TextHeadlessBlackNoteShortStem,
	"textTie": TextTie,
	"textTuplet3LongStem": TextTuplet3LongStem,
	"textTuplet3ShortStem": TextTuplet3ShortStem,
	"textTupletBracketEndLongStem": TextTupletBracketEndLongStem,
	"textTupletBracketEndShortStem": TextTupletBracketEndShortStem,
	"textTupletBracketStartLongStem": TextTupletBracketStartLongStem,
	"textTupletBracketStartShortStem": TextTupletBracketStartShortStem,
	"timeSig0": TimeSig0,
	"timeSig0Reversed": TimeSig0Reversed,
	"timeSig0Turned": TimeSig0Turned,
	"timeSig1": TimeSig1,
	"timeSig1Reversed": TimeSig1Reversed,
	"timeSig1Turned": TimeSig1Turned,
	"timeSig2": TimeSig2,
	"timeSig2Reversed": TimeSig2Reversed,
	"timeSig2Turned": TimeSig2Turned,
	"timeSig3": TimeSig3,
	"timeSig3Reversed": TimeSig3Reversed,
	"timeSig3Turned": TimeSig3Turned,
	"timeSig4": TimeSig4,
	"timeSig4Reversed": TimeSig4Reversed,
	"timeSig4Turned": TimeSig4Turned,
	"timeSig5": TimeSig5,
	"timeSig5Reversed": TimeSig5Reversed,
	"timeSig5Turned": TimeSig5Turned,
	"timeSig6": TimeSig6,
	"timeSig6Reversed": TimeSig6Reversed,
	"timeSig6Turned": TimeSig6Turned,
	"timeSig7": TimeSig7,
	"timeSig7Reversed": TimeSig7Reversed,
	"timeSig7Turned": TimeSig7Turned,
	"timeSig8": TimeSig8,
	"timeSig8Reversed": TimeSig8Reversed,
	"timeSig8Turned": TimeSig8Turned,
	"timeSig9": TimeSig9,
	"timeSig9Reversed": TimeSig9Reversed,
	"timeSig9Turned": TimeSig9Turned,
	"timeSigBracketLeft": TimeSigBracketLeft,
	"timeSigBracketLeftSmall": TimeSigBracketLeftSmall,
	"timeSigBracketRight": TimeSigBracketRight,
	"timeSigBracketRightSmall": TimeSigBracketRightSmall,
	"timeSigCombDenominator": TimeSigCombDenominator,
	"timeSigCombNumerator": TimeSigCombNumerator,
	"timeSigComma": TimeSigComma,
	"timeSigCommon": TimeSigCommon,
	"timeSigCommonReversed": TimeSigCommonReversed,
	"timeSigCommonTurned": TimeSigCommonTurned,
	"timeSigCut2": TimeSigCut2,
	"timeSigCut3": TimeSigCut3,
	"timeSigCutCommon": TimeSigCutCommon,
	"timeSigCutCommonReversed": TimeSigCutCommonReversed,
	"timeSigCutCommonTurned": TimeSigCutCommonTurned,
	"timeSigEquals": TimeSigEquals,
	"timeSigFractionHalf": TimeSigFractionHalf,
	"timeSigFractionOneThird": TimeSigFractionOneThird,
	"timeSigFractionQuarter": TimeSigFractionQuarter,
	"timeSigFractionThreeQuarters": TimeSigFractionThreeQuarters,
	"timeSigFractionTwoThirds": TimeSigFractionTwoThirds,
	"timeSigFractionalSlash": TimeSigFractionalSlash,
	"timeSigMinus": TimeSigMinus,
	"timeSigMultiply": TimeSigMultiply,
	"timeSigOpenPenderecki": TimeSigOpenPenderecki,
	"timeSigParensLeft": TimeSigParensLeft,
	"timeSigParensLeftSmall": TimeSigParensLeftSmall,
	"timeSigParensRight": TimeSigParensRight,
	"timeSigParensRightSmall": TimeSigParensRightSmall,
	"timeSigPlus": TimeSigPlus,
	"timeSigPlusSmall": TimeSigPlusSmall,
	"timeSigSlash": TimeSigSlash,
	"timeSigX": TimeSigX,
	"tremolo1": Tremolo1,
	"tremolo2": Tremolo2,
	"tremolo3": Tremolo3,
	"tremolo4": Tremolo4,
	"tremolo5": Tremolo5,
	"tremoloDivisiDots2": TremoloDivisiDots2,
	"tremoloDivisiDots3": TremoloDivisiDots3,
	"tremoloDivisiDots4": TremoloDivisiDots4,
	"tremoloDivisiDots6": TremoloDivisiDots6,
	"tremoloFingered1": TremoloFingered1,
	"tremoloFingered2": TremoloFingered2,
	"tremoloFingered3": TremoloFingered3,
	"tremoloFingered4": TremoloFingered4,
	"tremoloFingered5": TremoloFingered5,
	"tripleTongueAbove": TripleTongueAbove,
	"tripleTongueBelow": TripleTongueBelow,
	"tuplet0": Tuplet0,
	"tuplet1": Tuplet1,
	"tuplet2": Tuplet2,
	"tuplet3": Tuplet3,
	"tuplet4": Tuplet4,
	"tuplet5": Tuplet5,
	"tuplet6": Tuplet6,
	"tuplet7": Tuplet7,
	"tuplet8": Tuplet8,
	"tuplet9": Tuplet9,
	"tupletColon": TupletColon,
	"unmeasuredTremolo": UnmeasuredTremolo,
	"unmeasuredTremoloSimple": UnmeasuredTremoloSimple,
	"unpitchedPercussionClef1": UnpitchedPercussionClef1,
	"unpitchedPercussionClef2": UnpitchedPercussionClef2,
	"ventiduesima": Ventiduesima,
	"ventiduesimaAlta": VentiduesimaAlta,
	"ventiduesimaBassa": VentiduesimaBassa,
	"ventiduesimaBassaMb": VentiduesimaBassaMb,
	"vocalFingerClickStockhausen": VocalFingerClickStockhausen,
	"vocalHalbGesungen": VocalHalbGesungen,
	"vocalMouthClosed": VocalMouthClosed,
	"vocalMouthOpen": VocalMouthOpen,
	"vocalMouthPursed": VocalMouthPursed,
	"vocalMouthSlightlyOpen": VocalMouthSlightlyOpen,
	"vocalMouthWideOpen": VocalMouthWideOpen,
	"vocalNasalVoice": VocalNasalVoice,
	"vocalSprechgesang": VocalSprechgesang,
	"vocalTongueClickStockhausen": VocalTongueClickStockhausen,
	"vocalTongueFingerClickStockhausen": VocalTongueFingerClickStockhausen,
	"vocalsSussurando": VocalsSussurando,
	"wiggleArpeggiatoDown": WiggleArpeggiatoDown,
	"wiggleArpeggiatoDownArrow": WiggleArpeggiatoDownArrow,
	"wiggleArpeggiatoDownSwash": WiggleArpeggiatoDownSwash,
	"wiggleArpeggiatoUp": WiggleArpeggiatoUp,
	"wiggleArpeggiatoUpArrow": WiggleArpeggiatoUpArrow,
	"wiggleArpeggiatoUpSwash": WiggleArpeggiatoUpSwash,
	"wiggleCircular": WiggleCircular,
	"wiggleCircularConstant": WiggleCircularConstant,
	"wiggleCircularConstantFlipped": WiggleCircularConstantFlipped,
	"wiggleCircularConstantFlippedLarge": WiggleCircularConstantFlippedLarge,
	"wiggleCircularConstantLarge": WiggleCircularConstantLarge,
	"wiggleCircularEnd": WiggleCircularEnd,
	"wiggleCircularLarge": WiggleCircularLarge,
	"wiggleCircularLarger": WiggleCircularLarger,
	"wiggleCircularLargerStill": WiggleCircularLargerStill,
	"wiggleCircularLargest": WiggleCircularLargest,
	"wiggleCircularSmall": WiggleCircularSmall,
	"wiggleCircularStart": WiggleCircularStart,
	"wiggleGlissando": WiggleGlissando,
	"wiggleGlissandoGroup1": WiggleGlissandoGroup1,
	"wiggleGlissandoGroup2": WiggleGlissandoGroup2,
	"wiggleGlissandoGroup3": WiggleGlissandoGroup3,
	"wiggleRandom1": WiggleRandom1,
	"wiggleRandom2": WiggleRandom2,
	"wiggleRandom3": WiggleRandom3,
	"wiggleRandom4": WiggleRandom4,
	"wiggleSawtooth": WiggleSawtooth,
	"wiggleSawtoothNarrow": WiggleSawtoothNarrow,
	"wiggleSawtoothWide": WiggleSawtoothWide,
	"wiggleSquareWave": WiggleSquareWave,
	"wiggleSquareWaveNarrow": WiggleSquareWaveNarrow,
	"wiggleSquareWaveWide": WiggleSquareWaveWide,
	"wiggleTrill": WiggleTrill,
	"wiggleTrillFast": WiggleTrillFast,
	"wiggleTrillFaster": WiggleTrillFaster,
	"wiggleTrillFasterStill": WiggleTrillFasterStill,
	"wiggleTrillFastest": WiggleTrillFastest,
	"wiggleTrillSlow": WiggleTrillSlow,
	"wiggleTrillSlower": WiggleTrillSlower,
	"wiggleTrillSlowerStill": WiggleTrillSlowerStill,
	"wiggleTrillSlowest": WiggleTrillSlowest,
	"wiggleVIbratoLargestSlower": WiggleVIbratoLargestSlower,
	"wiggleVIbratoMediumSlower": WiggleVIbratoMediumSlower,
	"wiggleVibrato": WiggleVibrato,
	"wiggleVibratoLargeFast": WiggleVibratoLargeFast,
	"wiggleVibratoLargeFaster": WiggleVibratoLargeFaster,
	"wiggleVibratoLargeFasterStill": WiggleVibratoLargeFasterStill,
	"wiggleVibratoLargeFastest": WiggleVibratoLargeFastest,
	"wiggleVibratoLargeSlow": WiggleVibratoLargeSlow,
	"wiggleVibratoLargeSlower": WiggleVibratoLargeSlower,
	"wiggleVibratoLargeSlowest": WiggleVibratoLargeSlowest,
	"wiggleVibratoLargestFast": WiggleVibratoLargestFast,
	"wiggleVibratoLargestFaster": WiggleVibratoLargestFaster,
	"wiggleVibratoLargestFasterStill": WiggleVibratoLargestFasterStill,
	"wiggleVibratoLargestFastest": WiggleVibratoLargestFastest,
	"wiggleVibratoLargestSlow": WiggleVibratoLargestSlow,
	"wiggleVibratoLargestSlowest": WiggleVibratoLargestSlowest,
	"wiggleVibratoMediumFast": WiggleVibratoMediumFast,
	"wiggleVibratoMediumFaster": WiggleVibratoMediumFaster,
	"wiggleVibratoMediumFasterStill": WiggleVibratoMediumFasterStill,
	"wiggleVibratoMediumFastest": WiggleVibratoMediumFastest,
	"wiggleVibratoMediumSlow": WiggleVibratoMediumSlow,
	"wiggleVibratoMediumSlowest": WiggleVibratoMediumSlowest,
	"wiggleVibratoSmallFast": WiggleVibratoSmallFast,
	"wiggleVibratoSmallFaster": WiggleVibratoSmallFaster,
	"wiggleVibratoSmallFasterStill": WiggleVibratoSmallFasterStill,
	"wiggleVibratoSmallFastest": WiggleVibratoSmallFastest,
	"wiggleVibratoSmallSlow": WiggleVibratoSmallSlow,
	"wiggleVibratoSmallSlower": WiggleVibratoSmallSlower,
	"wiggleVibratoSmallSlowest": WiggleVibratoSmallSlowest,
	"wiggleVibratoSmallestFast": WiggleVibratoSmallestFast,
	"wiggleVibratoSmallestFaster": WiggleVibratoSmallestFaster,
	"wiggleVibratoSmallestFasterStill": WiggleVibratoSmallestFasterStill,
	"wiggleVibratoSmallestFastest": WiggleVibratoSmallestFastest,
	"wiggleVibratoSmallestSlow": WiggleVibratoSmallestSlow,
	"wiggleVibratoSmallestSlower": WiggleVibratoSmallestSlower,
	"wiggleVibratoSmallestSlowest": WiggleVibratoSmallestSlowest,
	"wiggleVibratoStart": WiggleVibratoStart,
	"wiggleVibratoWide": WiggleVibratoWide,
	"wiggleWavy": WiggleWavy,
	"wiggleWavyNarrow": WiggleWavyNarrow,
	"wiggleWavyWide": WiggleWavyWide,
	"windClosedHole": WindClosedHole,
	"windFlatEmbouchure": WindFlatEmbouchure,
	"windHalfClosedHole1": WindHalfClosedHole1,
	"windHalfClosedHole2": WindHalfClosedHole2,
	"windHalfClosedHole3": WindHalfClosedHole3,
	"windLessRelaxedEmbouchure": WindLessRelaxedEmbouchure,
	"windLessTightEmbouchure": WindLessTightEmbouchure,
	"windMouthpiecePop": WindMouthpiecePop,
	"windMultiphonicsBlackStem": WindMultiphonicsBlackStem,
	"windMultiphonicsBlackWhiteStem": WindMultiphonicsBlackWhiteStem,
	"windMultiphonicsWhiteStem": WindMultiphonicsWhiteStem,
	"windOpenHole": WindOpenHole,
	"windReedPositionIn": WindReedPositionIn,
	"windReedPositionNormal": WindReedPositionNormal,
	"windReedPositionOut": WindReedPositionOut,
	"windRelaxedEmbouchure": WindRelaxedEmbouchure,
	"windRimOnly": WindRimOnly,
	"windSharpEmbouchure": WindSharpEmbouchure,
	"windStrongAirPressure": WindStrongAirPressure,
	"windThreeQuartersClosedHole": WindThreeQuartersClosedHole,
	"windTightEmbouchure": WindTightEmbouchure,
	"windTrillKey": WindTrillKey,
	"windVeryTightEmbouchure": WindVeryTightEmbouchure,
	"windWeakAirPressure": WindWeakAirPressure,
}
func NameToCodepoint(name string) (rune, bool) {
	r, ok := nameToCodepoint[name]
	return r, ok
}
