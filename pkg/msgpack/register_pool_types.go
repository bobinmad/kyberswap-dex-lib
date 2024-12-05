package msgpack

// Code generated by github.com/KyberNetwork/kyberswap-dex-lib/pkg/msgpack/generate DO NOT EDIT.

import (
	"github.com/KyberNetwork/msgpack/v5"

	pkg_liquiditysource_algebra_integral "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/algebra/integral"
	pkg_liquiditysource_algebra_v1 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/algebra/v1"
	pkg_liquiditysource_balancerv1 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/balancer-v1"
	pkg_liquiditysource_balancerv2_composablestable "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/balancer-v2/composable-stable"
	pkg_liquiditysource_balancerv2_stable "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/balancer-v2/stable"
	pkg_liquiditysource_balancerv2_weighted "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/balancer-v2/weighted"
	pkg_liquiditysource_bancorv21 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/bancor-v21"
	pkg_liquiditysource_bancorv3 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/bancor-v3"
	pkg_liquiditysource_bebop "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/bebop"
	pkg_liquiditysource_bedrock_unieth "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/bedrock/unieth"
	pkg_liquiditysource_clipper "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/clipper"
	pkg_liquiditysource_curve_plain "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/curve/plain"
	pkg_liquiditysource_curve_stablemetang "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/curve/stable-meta-ng"
	pkg_liquiditysource_curve_stableng "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/curve/stable-ng"
	pkg_liquiditysource_curve_tricryptong "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/curve/tricrypto-ng"
	pkg_liquiditysource_curve_twocryptong "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/curve/twocrypto-ng"
	pkg_liquiditysource_daiusds "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/dai-usds"
	pkg_liquiditysource_deltaswapv1 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/deltaswap-v1"
	pkg_liquiditysource_dexalot "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/dexalot"
	pkg_liquiditysource_dodo_classical "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/dodo/classical"
	pkg_liquiditysource_dodo_dpp "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/dodo/dpp"
	pkg_liquiditysource_dodo_dsp "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/dodo/dsp"
	pkg_liquiditysource_dodo_dvm "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/dodo/dvm"
	pkg_liquiditysource_ethena_susde "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/ethena/susde"
	pkg_liquiditysource_ethervista "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/ether-vista"
	pkg_liquiditysource_etherfi_eeth "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/etherfi/eeth"
	pkg_liquiditysource_etherfi_vampire "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/etherfi/vampire"
	pkg_liquiditysource_etherfi_weeth "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/etherfi/weeth"
	pkg_liquiditysource_fluid_dext1 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/fluid/dex-t1"
	pkg_liquiditysource_fluid_vaultt1 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/fluid/vault-t1"
	pkg_liquiditysource_frax_sfrxeth "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/frax/sfrxeth"
	pkg_liquiditysource_frax_sfrxethconvertor "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/frax/sfrxeth-convertor"
	pkg_liquiditysource_genericsimplerate "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/generic-simple-rate"
	pkg_liquiditysource_gyroscope_2clp "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/gyroscope/2clp"
	pkg_liquiditysource_gyroscope_3clp "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/gyroscope/3clp"
	pkg_liquiditysource_gyroscope_eclp "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/gyroscope/eclp"
	pkg_liquiditysource_hashflowv3 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/hashflow-v3"
	pkg_liquiditysource_integral "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/integral"
	pkg_liquiditysource_kelp_rseth "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/kelp/rseth"
	pkg_liquiditysource_litepsm "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/litepsm"
	pkg_liquiditysource_maker_savingsdai "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/maker/savingsdai"
	pkg_liquiditysource_mantle_meth "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/mantle/meth"
	pkg_liquiditysource_mkrsky "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/mkr-sky"
	pkg_liquiditysource_mxtrading "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/mx-trading"
	pkg_liquiditysource_nativev1 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/native-v1"
	pkg_liquiditysource_nomiswap_nomiswapstable "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/nomiswap/nomiswapstable"
	pkg_liquiditysource_ondousdy "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/ondo-usdy"
	pkg_liquiditysource_primeeth "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/primeeth"
	pkg_liquiditysource_puffer_pufeth "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/puffer/pufeth"
	pkg_liquiditysource_renzo_ezeth "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/renzo/ezeth"
	pkg_liquiditysource_ringswap "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/ringswap"
	pkg_liquiditysource_rocketpool_reth "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/rocketpool/reth"
	pkg_liquiditysource_staderethx "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/staderethx"
	pkg_liquiditysource_swaapv2 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/swaap-v2"
	pkg_liquiditysource_swell_rsweth "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/swell/rsweth"
	pkg_liquiditysource_swell_sweth "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/swell/sweth"
	pkg_liquiditysource_uniswapv1 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/uniswap-v1"
	pkg_liquiditysource_uniswapv2 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/uniswap-v2"
	pkg_liquiditysource_usd0pp "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/usd0pp"
	pkg_liquiditysource_velocorev2_cpmm "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/velocore-v2/cpmm"
	pkg_liquiditysource_velocorev2_wombatstable "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/velocore-v2/wombat-stable"
	pkg_liquiditysource_velodromev1 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/velodrome-v1"
	pkg_liquiditysource_velodromev2 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/velodrome-v2"
	pkg_liquiditysource_woofiv2 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/woofi-v2"
	pkg_liquiditysource_woofiv21 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/liquidity-source/woofi-v21"
	pkg_source_algebrav1 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/algebrav1"
	pkg_source_balancer_stable "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/balancer/stable"
	pkg_source_balancer_weighted "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/balancer/weighted"
	pkg_source_balancercomposablestable "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/balancer-composable-stable"
	pkg_source_camelot "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/camelot"
	pkg_source_curve_aave "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/curve/aave"
	pkg_source_curve_base "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/curve/base"
	pkg_source_curve_compound "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/curve/compound"
	pkg_source_curve_meta "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/curve/meta"
	pkg_source_curve_plainoracle "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/curve/plain-oracle"
	pkg_source_curve_tricrypto "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/curve/tricrypto"
	pkg_source_curve_two "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/curve/two"
	pkg_source_dmm "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/dmm"
	pkg_source_dodo "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/dodo"
	pkg_source_elastic "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/elastic"
	pkg_source_equalizer "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/equalizer"
	pkg_source_fraxswap "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/fraxswap"
	pkg_source_fulcrom "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/fulcrom"
	pkg_source_fxdx "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/fxdx"
	pkg_source_gmx "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/gmx"
	pkg_source_gmxglp "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/gmx-glp"
	pkg_source_iziswap "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/iziswap"
	pkg_source_kokonutcrypto "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/kokonut-crypto"
	pkg_source_kyberpmm "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/kyber-pmm"
	pkg_source_lido "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/lido"
	pkg_source_lidosteth "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/lido-steth"
	pkg_source_limitorder "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/limitorder"
	pkg_source_liquiditybookv20 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/liquiditybookv20"
	pkg_source_liquiditybookv21 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/liquiditybookv21"
	pkg_source_madmex "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/madmex"
	pkg_source_makerpsm "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/makerpsm"
	pkg_source_mantisswap "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/mantisswap"
	pkg_source_maverickv1 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/maverickv1"
	pkg_source_metavault "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/metavault"
	pkg_source_nuriv2 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/nuriv2"
	pkg_source_pancakev3 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pancakev3"
	pkg_source_platypus "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/platypus"
	pkg_source_polmatic "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/pol-matic"
	pkg_source_quickperps "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/quickperps"
	pkg_source_ramsesv2 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/ramsesv2"
	pkg_source_saddle "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/saddle"
	pkg_source_slipstream "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/slipstream"
	pkg_source_smardex "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/smardex"
	pkg_source_solidlyv3 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/solidly-v3"
	pkg_source_swapbasedperp "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/swapbased-perp"
	pkg_source_syncswap_syncswapclassic "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/syncswap/syncswapclassic"
	pkg_source_syncswap_syncswapstable "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/syncswap/syncswapstable"
	pkg_source_synthetix "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/synthetix"
	pkg_source_uniswap "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/uniswap"
	pkg_source_uniswapv3 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/uniswapv3"
	pkg_source_usdfi "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/usdfi"
	pkg_source_velocimeter "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/velocimeter"
	pkg_source_velodrome "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/velodrome"
	pkg_source_velodromev2 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/velodromev2"
	pkg_source_vooi "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/vooi"
	pkg_source_wombat_wombatlsd "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/wombat/wombatlsd"
	pkg_source_wombat_wombatmain "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/wombat/wombatmain"
	pkg_source_woofiv2 "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/woofiv2"
	pkg_source_zkerafinance "github.com/KyberNetwork/kyberswap-dex-lib/pkg/source/zkera-finance"
)
func init() {
	msgpack.RegisterConcreteType(&pkg_liquiditysource_algebra_integral.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_algebra_v1.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_balancerv1.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_balancerv2_composablestable.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_balancerv2_stable.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_balancerv2_weighted.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_bancorv21.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_bancorv3.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_bebop.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_bedrock_unieth.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_clipper.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_curve_plain.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_curve_stablemetang.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_curve_stableng.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_curve_tricryptong.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_curve_twocryptong.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_daiusds.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_deltaswapv1.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_dexalot.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_dodo_classical.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_dodo_dpp.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_dodo_dsp.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_dodo_dvm.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_ethena_susde.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_ethervista.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_etherfi_eeth.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_etherfi_vampire.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_etherfi_weeth.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_fluid_dext1.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_fluid_vaultt1.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_frax_sfrxeth.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_frax_sfrxethconvertor.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_genericsimplerate.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_gyroscope_2clp.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_gyroscope_3clp.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_gyroscope_eclp.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_hashflowv3.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_integral.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_kelp_rseth.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_litepsm.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_maker_savingsdai.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_mantle_meth.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_mkrsky.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_mxtrading.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_nativev1.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_nomiswap_nomiswapstable.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_ondousdy.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_primeeth.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_puffer_pufeth.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_renzo_ezeth.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_ringswap.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_rocketpool_reth.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_staderethx.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_swaapv2.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_swell_rsweth.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_swell_sweth.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_uniswapv1.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_uniswapv2.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_usd0pp.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_velocorev2_cpmm.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_velocorev2_wombatstable.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_velodromev1.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_velodromev2.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_woofiv2.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_liquiditysource_woofiv21.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_algebrav1.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_balancer_stable.StablePool{})
	msgpack.RegisterConcreteType(&pkg_source_balancer_weighted.WeightedPool2Tokens{})
	msgpack.RegisterConcreteType(&pkg_source_balancercomposablestable.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_camelot.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_curve_aave.AavePool{})
	msgpack.RegisterConcreteType(&pkg_source_curve_base.PoolBaseSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_curve_compound.CompoundPool{})
	msgpack.RegisterConcreteType(&pkg_source_curve_meta.Pool{})
	msgpack.RegisterConcreteType(&pkg_source_curve_plainoracle.Pool{})
	msgpack.RegisterConcreteType(&pkg_source_curve_tricrypto.Pool{})
	msgpack.RegisterConcreteType(&pkg_source_curve_two.Pool{})
	msgpack.RegisterConcreteType(&pkg_source_dmm.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_dodo.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_elastic.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_equalizer.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_fraxswap.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_fulcrom.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_fxdx.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_gmx.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_gmxglp.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_iziswap.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_kokonutcrypto.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_kyberpmm.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_lido.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_lidosteth.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_limitorder.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_liquiditybookv20.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_liquiditybookv21.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_madmex.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_makerpsm.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_mantisswap.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_maverickv1.Pool{})
	msgpack.RegisterConcreteType(&pkg_source_metavault.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_nuriv2.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_pancakev3.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_platypus.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_polmatic.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_quickperps.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_ramsesv2.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_saddle.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_slipstream.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_smardex.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_solidlyv3.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_swapbasedperp.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_syncswap_syncswapclassic.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_syncswap_syncswapstable.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_synthetix.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_uniswap.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_uniswapv3.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_usdfi.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_velocimeter.Pool{})
	msgpack.RegisterConcreteType(&pkg_source_velodrome.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_velodromev2.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_vooi.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_wombat_wombatlsd.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_wombat_wombatmain.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_woofiv2.PoolSimulator{})
	msgpack.RegisterConcreteType(&pkg_source_zkerafinance.PoolSimulator{})
}
